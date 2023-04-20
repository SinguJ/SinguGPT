package access

// Event 事件
type Event string

// UserChangeEvent 用户变更事件
type UserChangeEvent Event

const (
    // EventCreateUser 创建用户事件
    EventCreateUser UserChangeEvent = "create-user"
    // EventModifyUser 修改用户事件
    EventModifyUser = "modify-user"
    // EventDeleteUser 删除用户事件
    EventDeleteUser = "delete-user"
)
