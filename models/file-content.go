package models

import (
    "fmt"
    "io"
    "os"
    "path"
)

// 内容类型与文件扩展名映射
var contentTypeExtNameMapping = map[ContentType][]string{
    ContentTypeHTML:     {".html"},
    ContentTypeMarkdown: {".md", ".markdown"},
    ContentTypeText: {
        ".c",
        ".h",
        ".cpp",
        ".go",
        ".java",
        ".js",
        ".py",
        ".json",
        ".yaml",
        ".text",
    },
    ContentTypeBytes: {
        ".exe",
        ".so",
        ".dll",
        ".tar",
        ".gz",
        ".zip",
        ".rar",
        ".7z",
        ".iso",
        ".dmg",
        ".rpm",
        ".deb",
    },
}

// 根据文件后缀名匹配内容类型
func matchContentType(extName string) ContentType {
    for contentType, mapping := range contentTypeExtNameMapping {
        for _, _extName := range mapping {
            if _extName == extName {
                return contentType
            }
        }
    }
    return ContentTypeUnknown
}

// FileContent 文件类型的内容
type FileContent struct {
    // 文件所在目录的路径
    dirPath string
    // 文件名称
    // 不包含文件后缀的名称
    name string
    // 文件扩展名
    // 包含 '.' 前缀
    extName string
    // 文件内容字节数
    size int64
    // 文件内容
    // 若 content 不为 nil，则表示该文件是一个存储在内存中的虚拟文件
    content Content
}

// Type 内容的类型
func (f *FileContent) Type() ContentType {
    return ContentTypeFile
}

// FileType 文件类型
func (f *FileContent) FileType() ContentType {
    if f.content != nil {
        return f.content.Type()
    }
    return matchContentType(f.extName)
}

// Tag 内容标记
func (f *FileContent) Tag() Tag {
    return TagFile
}

// Path 文件路径
func (f *FileContent) Path() string {
    return fmt.Sprintf("%s/%s%s", f.dirPath, f.name, f.extName)
}

// DirPath 文件所在目录的路径
func (f *FileContent) DirPath() string {
    return f.dirPath
}

// Name 文件名称
// 不包含文件扩展名
func (f *FileContent) Name() string {
    return f.name
}

// FullName 完整文件名称
// 包含文件扩展名
func (f *FileContent) FullName() string {
    return fmt.Sprintf("%s%s", f.name, f.extName)
}

// ExtName 文件扩展名
// 包含 '.' 前缀
func (f *FileContent) ExtName() string {
    return f.extName
}

// Len 内容长度
func (f *FileContent) Len() int64 {
    return f.size
}

// ToBytes 转为字节数组
func (f *FileContent) ToBytes() []byte {
    if f.content != nil {
        return f.content.ToBytes()
    }
    byteArr, err := os.ReadFile(f.Path())
    if err != nil {
        panic(err)
    }
    return byteArr
}

// ToString 转为字符串
func (f *FileContent) ToString() string {
    if f.content != nil {
        return f.content.ToString()
    }
    return string(f.ToBytes())
}

// ToReader 转为字节读取流
func (f *FileContent) ToReader() io.Reader {
    if f.content != nil {
        return f.content.ToReader()
    }
    file, err := os.Open(f.Path())
    if err != nil {
        panic(err)
    }
    return file
}

// 拆分文件路径
func filepathSplit(filepath string) (dirPath string, name string, extName string) {
    dirPath = path.Dir(filepath)
    fullName := path.Base(filepath)
    extName = path.Ext(fullName)
    name = fullName[:len(fullName)-len(extName)]
    return
}

// 获取文件字节数
func getFileSize(filepath string) int64 {
    file, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    defer func(file *os.File) {
        err := file.Close()
        if err != nil {
            panic(err)
        }
    }(file)

    // 获取文件信息
    fileInfo, err := file.Stat()
    if err != nil {
        panic(err)
    }

    // 获取文件字节数
    return fileInfo.Size()
}

// 创建  FileContent 对象
func newFileContent(dirPath string, name string, extName string, content Content) *FileContent {
    var size int64
    if content != nil {
        size = int64(len(content.ToBytes()))
    } else {
        size = getFileSize(fmt.Sprintf("%s/%s%s", dirPath, name, extName))
    }
    return &FileContent{
        dirPath: dirPath,
        name:    name,
        extName: extName,
        size:    size,
        content: content,
    }
}

// NewFileContent 将指定的 Content 包装成 FileContent
func NewFileContent(filepath string, content Content) *FileContent {
    dirPath, name, extName := filepathSplit(filepath)
    if fc, ok := content.(*FileContent); ok {
        content = fc.content
    }
    return newFileContent(dirPath, name, extName, content)
}

// NewFileContentByLocalFile 创建本地文件的 FileContent 对象
func NewFileContentByLocalFile(filepath string) *FileContent {
    dirPath, name, extName := filepathSplit(filepath)
    return NewFileContentByLocalFile2(dirPath, name, extName)
}

// NewFileContentByLocalFile2 创建本地文件的 FileContent 对象
func NewFileContentByLocalFile2(dirPath string, name string, extName string) *FileContent {
    return newFileContent(dirPath, name, extName, nil)
}
