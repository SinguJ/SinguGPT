package imap

import (
    "bytes"
    "fmt"
    "io"
    "strings"

    "github.com/emersion/go-imap/v2"
    "github.com/emersion/go-imap/v2/imapclient"
    _ "github.com/emersion/go-message/charset"
    mailLib "github.com/emersion/go-message/mail"

    "SinguGPT/errors"
)

// 读取地址信息
func readAddress(addrArr []imap.Address) [][2]string {
    arr := make([][2]string, len(addrArr))
    for index, addr := range addrArr {
        arr[index] = [2]string{addr.Addr(), addr.Name}
    }
    return arr
}

func getBody(msg *imapclient.FetchMessageBuffer) []byte {
    for _, literal := range msg.BodySection {
        if literal != nil {
            return literal
        }
    }
    return nil
}

func readBody(msg *imapclient.FetchMessageBuffer) ([]*Content, []*Attach) {
    bodyStructure := msg.BodyStructure.(*imap.BodyStructureMultiPart)
    boundary := bodyStructure.Extended.Params["boundary"]
    //goland:noinspection SpellCheckingInspection
    tag := "IMTHEBOUNDARY"
    mediaType := strings.TrimSpace(bodyStructure.MediaType())
    prefix := fmt.Sprintf("Mime-Version: 1.0\r\nContent-Type: %s; boundary=%s\r\n\r\n", mediaType, tag)
    text := string(getBody(msg))
    text = prefix + strings.ReplaceAll(text, boundary, tag)
    reader, err := mailLib.CreateReader(bytes.NewReader([]byte(text)))
    if err != nil {
        panic(errors.Wrap(err))
    }
    contents := make([]*Content, 0)
    attaches := make([]*Attach, 0)
    var part *mailLib.Part
    for {
        part, err = reader.NextPart()
        if err != nil {
            if err == io.EOF {
                break
            }
            panic(errors.Wrap(err))
        }
        switch part.Header.(type) {
        case *mailLib.InlineHeader:
            inlineHeader := part.Header.(*mailLib.InlineHeader)
            // 构建内容对象
            content := Content{}
            // 读取内容类型
            contentType, _, err := inlineHeader.ContentType()
            if err != nil {
                panic(errors.Wrap(err))
            }
            switch contentType {
            case "text/plain":
                content.Type = Text
            case "text/html":
                content.Type = HTML
            default:
                content.Type = Other
            }
            // 读取内容长度
            content.Len = inlineHeader.Len()
            // 读取内容文本
            data, err := io.ReadAll(part.Body)
            if err != nil {
                panic(errors.Wrap(err))
            }
            content.Text = string(data)
            contents = append(contents, &content)
        case *mailLib.AttachmentHeader:
            attachmentHeader := part.Header.(*mailLib.AttachmentHeader)
            // 构建附件对象
            attach := Attach{}
            // 读取内容类型
            contentType, _, err := attachmentHeader.ContentType()
            if err != nil {
                panic(errors.Wrap(err))
            }
            switch contentType {
            // TODO: 补充文件的 ContentType
            default:
                attach.Type = Other
            }
            // 读取文件名
            filename, err := attachmentHeader.Filename()
            if err != nil {
                panic(errors.Wrap(err))
            }
            attach.Filename = filename
            // 读取文件数据&文件大小
            data, err := io.ReadAll(part.Body)
            if err != nil {
                panic(errors.Wrap(err))
            }
            attach.Bytes = data
            attach.Size = len(data)
            attaches = append(attaches, &attach)
        }
    }
    return contents, attaches
}

// 读取信件
func readMail(msg *imapclient.FetchMessageBuffer) *Mail {
    // 创建邮件对象
    mail := Mail{
        ID:      msg.Envelope.MessageID,
        SeqNum:  msg.SeqNum,
        From:    readAddress(msg.Envelope.From),
        To:      readAddress(msg.Envelope.To),
        Date:    msg.InternalDate,
        Subject: msg.Envelope.Subject,
    }
    // 读取正文及附件
    mail.Contents, mail.Attaches = readBody(msg)
    return &mail
}
