package cronjob

import (
	"time"
)

// MessagePubMetadata 定义消息数据
type MessagePubMetadata struct {
	Code              int                    `json:"code"`                          // 状态
	RefType           string                 `json:"ref_type"`                      // 关联消息类型
	RefID             string                 `json:"ref_id"`                        // 关联消息id
	Title             string                 `json:"title,omitempty"`               // 标题
	Subtitle          string                 `json:"subtitle,omitempty"`            // 子标题
	MessageType       int                    `json:"message_type,omitempty"`        // 消息类型
	Message           string                 `json:"message,omitempty"`             // 消息
	CreatorID         string                 `json:"creator_id"`                    // 创建人id
	NotifyTo          []string               `json:"notify_to"`                     // 推送用户数组
	AffectedUID       []string               `json:"affected_uid"`                  // 受影响用户
	StartAt           int64                  `json:"start_at,omitempty"`            // 开始时间
	EndAt             int64                  `json:"end_at,omitempty"`              // 结束时间
	RemindAt          int64                  `json:"remind_at,omitempty"`           // 提醒时间
	CompleteAt        int64                  `json:"complete_at,omitempty"`         // 完成时间
	CancelAt          int64                  `json:"cancel_at,omitempty"`           // 取消时间
	CreateAt          int64                  `json:"create_at,omitempty"`           // 创建时间
	LinkURL           string                 `json:"link_url,omitempty"`            // 点击跳转url
	CommentID         string                 `json:"comment_id,omitempty"`          // 评论ID
	AlreadyRead       bool                   `json:"already_read,omitempty"`        // 已读
	IsRedDot          bool                   `json:"is_red_dot,omitempty"`          // 是否红点
	FileID            string                 `json:"file_id,omitempty"`             // 文件id
	NotStore          bool                   `json:"not_store,omitempty"`           // 不存储到数据库
	RepeatID          string                 `json:"repeat_id,omitempty"`           // 重复事项重复id
	MsgType           int                    `json:"msg_type,omitempty"`            // 消息类型
	MsgFormat         int                    `json:"msg_format,omitempty"`          // 消息格式
	SystemType        int                    `json:"system_type,omitempty"`         // 系统消息
	Changes           map[string]interface{} `json:"changes,omitempty"`             // 变更信息
	NotPushSocket     bool                   `json:"not_push_socket,omitempty"`     // 不推送socket
	BatchID           string                 `json:"batch_id,omitempty"`            // 是否批量操作
	BatchType         int                    `json:"batch_type,omitempty"`          // 批量类型
	BatchRefID        string                 `json:"batch_ref_id,omitempty"`        // 批量操作关联id
	WorkspaceNoticeID string                 `json:"workspace_notice_id,omitempty"` // 空间通知id
	KrID              string                 `json:"kr_id,omitempty"`               // 目标KR ID
	Header            MeaasgeHeader          `json:"header,omitempty"`              // 消息header
}

// MeaasgeHeader 消息header
type MeaasgeHeader struct {
	CurrentVersion string            `json:"current_version,omitempty"` // 当前版本
	MaxVersion     string            `json:"max_version,omitempty"`     // 最大版本
	MinVersion     string            `json:"min_version,omitempty"`     // 最小版本
	SendForm       string            `json:"send_form,omitempty"`       // 发送平台
	Fn             []MessageHeaderFn `json:"fn,omitempty"`              // 功能
}

// setFn 设置功能
func (m *MeaasgeHeader) SetFn(f MessageHeaderFn) {
	m.Fn = append(m.Fn, f)
}

// MessageHeaderFn 消息header功能
type MessageHeaderFn struct {
	Type       int      `json:"type,omitempty"`        // 功能类型，1:红点，2：角标，3：弹窗，4：新邀请标记
	UserID     []string `json:"user_id,omitempty"`     // 用户id
	MaxVersion string   `json:"max_version,omitempty"` // 最大版本
	MinVersion string   `json:"min_version,omitempty"` // 最小版本
}

// NewMessageData 实例化
func NewMessageData() *MessagePubMetadata {
	return &MessagePubMetadata{
		NotifyTo: make([]string, 0),
		CreateAt: time.Now().Unix(),
	}
}

// SetCode 设置code
func (m *MessagePubMetadata) SetCode(code int) *MessagePubMetadata {
	m.Code = code
	return m
}

// SetRefType 设置关联类型 task、record
func (m *MessagePubMetadata) SetRefType(refType string) *MessagePubMetadata {
	m.RefType = refType
	return m
}

// SetRefID 设置关联id
func (m *MessagePubMetadata) SetRefID(refID string) *MessagePubMetadata {
	m.RefID = refID
	return m
}

// SetTitle 设置标题
func (m *MessagePubMetadata) SetTitle(title string) *MessagePubMetadata {
	m.Title = title
	return m
}

// SetMessage 设置消息
func (m *MessagePubMetadata) SetMessage(message string) *MessagePubMetadata {
	m.Message = message
	return m
}

// SetCreatorID 设置创建人id
func (m *MessagePubMetadata) SetCreatorID(creatorID string) *MessagePubMetadata {
	m.CreatorID = creatorID
	return m
}

// SetNotifyTo 设置通知用户
func (m *MessagePubMetadata) SetNotifyTo(notifyTo []string) *MessagePubMetadata {
	m.NotifyTo = notifyTo
	return m
}

// SetAffectedUID 设置受影响人
func (m *MessagePubMetadata) SetAffectedUID(affectedUID []string) *MessagePubMetadata {
	m.AffectedUID = affectedUID
	return m
}

// SetStartAt 设置开始时间
func (m *MessagePubMetadata) SetStartAt(startAt int64) *MessagePubMetadata {
	m.StartAt = startAt
	return m
}

// SetEndAt 设置结束时间
func (m *MessagePubMetadata) SetEndAt(endAt int64) *MessagePubMetadata {
	m.EndAt = endAt
	return m
}

// SetRemindAt 设置提醒时间
func (m *MessagePubMetadata) SetRemindAt(remindAt int64) *MessagePubMetadata {
	m.RemindAt = remindAt
	return m
}

// SetCompleteAt 设置完成时间
func (m *MessagePubMetadata) SetCompleteAt(completeAt int64) *MessagePubMetadata {
	m.CompleteAt = completeAt
	return m
}

// SetCancelAt 设置取消时间
func (m *MessagePubMetadata) SetCancelAt(cancelAt int64) *MessagePubMetadata {
	m.CancelAt = cancelAt
	return m
}

// SetCreateAt 设置创建时间
func (m *MessagePubMetadata) SetCreateAt(createAt int64) *MessagePubMetadata {
	m.CreateAt = createAt
	return m
}

// SetLinkURL 设置跳转url
func (m *MessagePubMetadata) SetLinkURL(linkURL string) *MessagePubMetadata {
	m.LinkURL = linkURL
	return m
}

// SetFileID 设置文件id
func (m *MessagePubMetadata) SetFileID(fileID string) *MessagePubMetadata {
	m.FileID = fileID
	return m
}

// SetNotStore 设置不存储数据库
func (m *MessagePubMetadata) SetNotStore() *MessagePubMetadata {
	m.NotStore = true
	return m
}

// SetRepeatID 设置任务重复id
func (m *MessagePubMetadata) SetRepeatID(repeatID string) *MessagePubMetadata {
	m.RepeatID = repeatID
	return m
}

// SetMsgType 设置消息类型
func (m *MessagePubMetadata) SetMsgType(msgType int) *MessagePubMetadata {
	m.MsgType = msgType
	return m
}

// SetMsgFormat 设置消息格式
func (m *MessagePubMetadata) SetMsgFormat(msgForamt int) *MessagePubMetadata {
	m.MsgFormat = msgForamt
	return m
}

// SetSystemType 设置系统消息类型
func (m *MessagePubMetadata) SetSystemType(systemType int) *MessagePubMetadata {
	m.SystemType = systemType
	return m
}

// SetChanges 设置变更信息
func (m *MessagePubMetadata) SetChanges(changes map[string]interface{}) *MessagePubMetadata {
	m.Changes = changes
	return m
}

// SetNotPushSocket 设置不褪色socket
func (m *MessagePubMetadata) SetNotPushSocket(notPushSocket bool) *MessagePubMetadata {
	m.NotPushSocket = notPushSocket
	return m
}

// SetBatchID 设置批量id
func (m *MessagePubMetadata) SetBatchID(batchID string) *MessagePubMetadata {
	m.BatchID = batchID
	return m
}

// SetBatchType 设置批量类型
func (m *MessagePubMetadata) SetBatchType(batchType int) *MessagePubMetadata {
	m.BatchType = batchType
	return m
}

// SetBatchRefID 设置批量操作关联id
func (m *MessagePubMetadata) SetBatchRefID(batchRefID string) *MessagePubMetadata {
	m.BatchRefID = batchRefID
	return m
}

// SetWorkspaceNoticeID 设置空间通知id
func (m *MessagePubMetadata) SetWorkspaceNoticeID(workspaceNoticeID string) *MessagePubMetadata {
	m.WorkspaceNoticeID = workspaceNoticeID
	return m
}

// SetMessageHeader 设置消息功能header
func (m *MessagePubMetadata) SetMessageHeader(header MeaasgeHeader) *MessagePubMetadata {
	m.Header = header
	return m
}
