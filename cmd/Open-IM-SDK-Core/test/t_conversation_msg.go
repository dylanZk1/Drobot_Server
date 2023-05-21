package test

import (
	"encoding/json"
	"fmt"
	"open_im_sdk/internal/login"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/open_im_sdk_callback"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
	"open_im_sdk/sdk_struct"
	"sync"
)

//func DotestSetConversationRecvMessageOpt() {
//	var callback BaseSuccessFailedTest
//	callback.funcName = utils.GetSelfFuncName()
//	var idList []string
//	idList = append(idList, "18567155635")
//	jsontest, _ := json.Marshal(idList)
//	open_im_sdk.SetConversationRecvMessageOpt(&callback, string(jsontest), 2)
//	fmt.Println("SetConversationRecvMessageOpt", string(jsontest))
//}
//
//func DoTestGetMultipleConversation() {
//	var callback BaseSuccessFailedTest
//	callback.funcName = utils.GetSelfFuncName()
//	var idList []string
//	fmt.Println("DoTestGetMultipleConversation come here")
//	idList = append(idList, "single_13977954313", "group_77215e1b13b75f3ab00cb6570e3d9618")
//	jsontest, _ := json.Marshal(idList)
//	open_im_sdk.GetMultipleConversation(string(jsontest), &callback)
//	fmt.Println("GetMultipleConversation", string(jsontest))
//}
//
//func DoTestGetConversationRecvMessageOpt() {
//	var callback BaseSuccessFailedTest
//	callback.funcName = utils.GetSelfFuncName()
//	var idList []string
//	idList = append(idList, "18567155635")
//	jsontest, _ := json.Marshal(idList)
//	open_im_sdk.GetConversationRecvMessageOpt(&callback, string(jsontest))
//	fmt.Println("GetConversationRecvMessageOpt", string(jsontest))
//}

func DoTestDeleteAllMsgFromLocalAndSvr() {
	var deleteConversationCallback DeleteConversationCallBack
	operationID := utils.OperationIDGenerator()
	log.Info(operationID, utils.GetSelfFuncName(), "args ")
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(deleteConversationCallback, operationID)
}
func DoTestSearchLocalMessages() {
	//[SearchLocalMessages args:  {"conversationID":"single_707010937","keywordList":["1"],"keywordListMatchType":0,"senderUserIDList":[],"messageTypeList":[],"searchTimePosition":0,"searchTimePeriod":0,"pageIndex":1,"count":200}]
	var testSearchLocalMessagesCallBack SearchLocalMessagesCallBack
	testSearchLocalMessagesCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.SearchLocalMessagesParams
	params.KeywordList = []string{"1"}
	params.ConversationID = "super_group_3907826375"
	params.Count = 20
	params.PageIndex = 1
	//s:=strings.Trim(params.KeywordList[0],"")
	//fmt.Println(len(s),s)
	//params.KeywordListMatchType = 1
	params.MessageTypeList = []int{101, 106}
	open_im_sdk.SearchLocalMessages(testSearchLocalMessagesCallBack, testSearchLocalMessagesCallBack.OperationID, utils.StructToJsonString(params))
}

func DoTestGetHistoryMessage(userID string) {
	var testGetHistoryCallBack GetHistoryCallBack
	testGetHistoryCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.GetHistoryMessageListParams
	params.UserID = userID
	params.ConversationID = "super_group_3907826375"
	//params.StartClientMsgID = "97f12899778823019f13ea46b0c1e6dd"
	params.Count = 10
	open_im_sdk.GetHistoryMessageList(testGetHistoryCallBack, testGetHistoryCallBack.OperationID, utils.StructToJsonString(params))
}
func DoTestFindMessageList() {
	var testFindMessageListCallBack FindMessageListCallBack
	testFindMessageListCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.FindMessageListParams
	temp := sdk_params_callback.ConversationArgs{ConversationID: "super_group_4205679980", ClientMsgIDList: []string{"eee68d85a43991d6b2e7354c52c5321d", "736f40f902046a6e879dc7257d3e81df"}}
	//temp1 := sdk_params_callback.ConversationArgs{ConversationID: "super_group_3320742908", ClientMsgIDList: []string{"acf09fcdda48bf2cb39faba31ac63b5c", "b121d3a7f269636afd255b6001d3fc80", "d8951d1c5192ad39f37f44de93a83302"}}
	params = append(params, &temp)
	//params = append(params, &temp1)
	open_im_sdk.FindMessageList(testFindMessageListCallBack, testFindMessageListCallBack.OperationID, utils.StructToJsonString(params))
}
func DoTestSetMessageReactionExtensions() {
	var testSetMessageReactionExtensionsCallBack SetMessageReactionExtensionsCallBack
	testSetMessageReactionExtensionsCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.SetMessageReactionExtensionsParams
	var data server_api_params.KeyValue
	data.TypeKey = "7788"
	m := make(map[string]interface{})
	m["test1"] = "hello1"
	m["test2"] = "hello2"
	data.Value = utils.StructToJsonString(m)
	params = append(params, &data)
	s := sdk_struct.MsgStruct{}
	s.SessionType = 3
	s.GroupID = "1420026997"
	s.ClientMsgID = "831c270ae1d7472dc633e7be06b37db5"
	//params = append(params, &temp1)
	open_im_sdk.SetMessageReactionExtensions(testSetMessageReactionExtensionsCallBack, testSetMessageReactionExtensionsCallBack.OperationID, utils.StructToJsonString(s),
		utils.StructToJsonString(params))
}
func DoTestUpdateFcmToken() {
	var testUpdateFcmTokenCallBack UpdateFcmTokenCallBack
	testUpdateFcmTokenCallBack.OperationID = utils.OperationIDGenerator()
	open_im_sdk.UpdateFcmToken(testUpdateFcmTokenCallBack, "2132323", testUpdateFcmTokenCallBack.OperationID)
}
func DoTestSetAppBadge() {
	var testSetAppBadgeCallBack SetAppBadgeCallBack
	testSetAppBadgeCallBack.OperationID = utils.OperationIDGenerator()
	open_im_sdk.SetAppBadge(testSetAppBadgeCallBack, 100, testSetAppBadgeCallBack.OperationID)
}

func DoTestGetAdvancedHistoryMessageList() {
	var testGetHistoryCallBack GetHistoryCallBack
	var params sdk_params_callback.GetAdvancedHistoryMessageListParams
	params.UserID = ""
	params.ConversationID = "super_group_812146266"
	params.StartClientMsgID = "83ca933d559d0374258550dd656a661c"
	params.Count = 30
	//params.LastMinSeq = seq
	open_im_sdk.GetAdvancedHistoryMessageList(testGetHistoryCallBack, testGetHistoryCallBack.OperationID, utils.StructToJsonString(params))
}
func DoTestGetHistoryMessageReverse(userID string) {
	var testGetHistoryReverseCallBack GetHistoryReverseCallBack
	testGetHistoryReverseCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.GetHistoryMessageListParams
	params.UserID = userID
	params.Count = 10
	params.ConversationID = "single_707008149"
	params.StartClientMsgID = "d40dde77f29b14d3a16ca6f422776890"
	open_im_sdk.GetHistoryMessageListReverse(testGetHistoryReverseCallBack, testGetHistoryReverseCallBack.OperationID, utils.StructToJsonString(params))
}
func DoTestGetGroupHistoryMessage() {
	var testGetHistoryCallBack GetHistoryCallBack
	testGetHistoryCallBack.OperationID = utils.OperationIDGenerator()
	var params sdk_params_callback.GetHistoryMessageListParams
	params.GroupID = "cb7aaa8e5f83d92db2ed1573cd01870c"
	params.Count = 10
	open_im_sdk.GetHistoryMessageList(testGetHistoryCallBack, testGetHistoryCallBack.OperationID, utils.StructToJsonString(params))
}

//func DoTestGetGroupHistoryMessage() {
//	var testGetHistoryCallBack GetHistoryCallBack
//	testGetHistoryCallBack.OperationID = utils.OperationIDGenerator()
//	var params sdk_params_callback.GetHistoryMessageListParams
//	params.GroupID = "cb7aaa8e5f83d92db2ed1573cd01870c"
//	params.Count = 10
//	open_im_sdk.GetHistoryMessageList(testGetHistoryCallBack, testGetHistoryCallBack.OperationID, utils.StructToJsonString(params))
//}

//func DoTestDeleteConversation(conversationID string) {
//	var testDeleteConversation DeleteConversationCallBack
//	open_im_sdk.DeleteConversation(conversationID, testDeleteConversation)
//
//}

type DeleteConversationCallBack struct {
}

func (d DeleteConversationCallBack) OnError(errCode int32, errMsg string) {
	fmt.Printf("DeleteConversationCallBack , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (d DeleteConversationCallBack) OnSuccess(data string) {
	fmt.Printf("DeleteConversationCallBack , success,data:%v\n", data)
}

type DeleteMessageCallBack struct {
	Msg string
}

func (d DeleteMessageCallBack) OnError(errCode int32, errMsg string) {
	fmt.Printf("DeleteMessageCallBack , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (d *DeleteMessageCallBack) OnSuccess(data string) {
	fmt.Printf("DeleteMessageCallBack , success,data:%v\n", data)
	d.Msg = data
}

func (d DeleteMessageCallBack) GetMessage() string {
	return d.Msg
}

func DoTestDeleteMessageFromLocalAndSvr(callback open_im_sdk_callback.Base, message string) {
	cb := &DeleteMessageCallBack{}
	msg := server_api_params.MsgData{
		SendID:           "",
		RecvID:           "",
		GroupID:          "",
		ClientMsgID:      "",
		ServerMsgID:      "",
		SenderPlatformID: 0,
		SenderNickname:   "",
		SenderFaceURL:    "",
		SessionType:      0,
		MsgFrom:          0,
		ContentType:      0,
		Content:          nil,
		Seq:              0,
		SendTime:         0,
		CreateTime:       0,
		Status:           0,
		Options:          nil,
		OfflinePushInfo:  nil,
		//XXX_NoUnkeyedLiteral: struct{}{},
		//XXX_unrecognized:     nil,
		//XXX_sizecache:        0,
	}
	operationID := utils.OperationIDGenerator()
	open_im_sdk.DeleteMessageFromLocalAndSvr(cb, operationID, utils.StructToJsonString(msg))
}

func DoTestDeleteConversationMsgFromLocalAndSvr(conversationID string) {
	cb := &DeleteMessageCallBack{}
	operationID := utils.OperationIDGenerator()
	open_im_sdk.DeleteConversationFromLocalAndSvr(cb, operationID, conversationID)
}

type TestGetAllConversationListCallBack struct {
	OperationID string
}

func (t TestGetAllConversationListCallBack) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestGetAllConversationListCallBack ", errCode, errMsg)
}

func (t TestGetAllConversationListCallBack) OnSuccess(data string) {
	log.Info(t.OperationID, "ConversationCallBack ", data)
}

func DoTestGetAllConversation() {
	var test TestGetAllConversationListCallBack
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetAllConversationList(test, test.OperationID)
}

func DoTestGetOneConversation(friendID string) {
	var test TestGetAllConversationListCallBack
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetOneConversation(test, test.OperationID, constant.SingleChatType, friendID)
}

func DoTestGetConversations(conversationIDs string) {
	var test TestGetAllConversationListCallBack
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetMultipleConversation(test, test.OperationID, conversationIDs)
}

type TestSetConversationPinnedCallback struct {
	OperationID string
}

func (t TestSetConversationPinnedCallback) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestSetConversationPinnedCallback ", errCode, errMsg)
}

func (t TestSetConversationPinnedCallback) OnSuccess(data string) {
	log.Info(t.OperationID, "TestSetConversationPinnedCallback ", data)
}

func DoTestSetConversationRecvMessageOpt(conversationIDs []string, opt int) {
	var callback testProcessGroupApplication
	callback.OperationID = utils.OperationIDGenerator()
	log.Info(callback.OperationID, utils.GetSelfFuncName(), "input: ")
	s := utils.StructToJsonString(conversationIDs)
	open_im_sdk.SetConversationRecvMessageOpt(callback, callback.OperationID, s, opt)
}

func DoTestSetConversationPinned(conversationID string, pin bool) {
	var test TestSetConversationPinnedCallback
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.PinConversation(test, test.OperationID, conversationID, pin)
}

func DoTestSetOneConversationPrivateChat(conversationID string, privateChat bool) {
	var test TestSetConversationPinnedCallback
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.SetOneConversationPrivateChat(test, test.OperationID, conversationID, privateChat)
}

func DoTestSetBurnDuration(conversationID string) {
	var test TestSetConversationPinnedCallback
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.SetOneConversationBurnDuration(test, test.OperationID, conversationID, 180)
}

func DoTestSetOneConversationRecvMessageOpt(conversationID string, opt int) {
	var test TestSetConversationPinnedCallback
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.SetOneConversationRecvMessageOpt(test, test.OperationID, conversationID, opt)
}

type TestGetConversationListSplitCallBack struct {
	OperationID string
}

func (t TestGetConversationListSplitCallBack) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestGetConversationListSplitCallBack err ", errCode, errMsg)
}

func (t TestGetConversationListSplitCallBack) OnSuccess(data string) {
	log.Info(t.OperationID, "TestGetConversationListSplitCallBack  success", data)
}
func DoTestGetConversationListSplit() {
	var test TestGetConversationListSplitCallBack
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetConversationListSplit(test, test.OperationID, 1, 2)
}

type TestGetOneConversationCallBack struct {
}

func (t TestGetOneConversationCallBack) OnError(errCode int32, errMsg string) {
	fmt.Printf("TestGetOneConversationCallBack , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (t TestGetOneConversationCallBack) OnSuccess(data string) {
	fmt.Printf("TestGetOneConversationCallBack , success,data:%v\n", data)
}

func DoTestGetConversationRecvMessageOpt(list string) {
	var test TestGetConversationRecvMessageOpt
	test.OperationID = utils.OperationIDGenerator()
	open_im_sdk.GetConversationRecvMessageOpt(test, test.OperationID, list)
}

type TestGetConversationRecvMessageOpt struct {
	OperationID string
}

func (t TestGetConversationRecvMessageOpt) OnError(errCode int32, errMsg string) {
	fmt.Printf("TestGetConversationRecvMessageOpt , errCode:%v,errMsg:%v\n", errCode, errMsg)
}

func (t TestGetConversationRecvMessageOpt) OnSuccess(data string) {
	fmt.Printf("TestGetConversationRecvMessageOpt , success,data:%v\n", data)
}

//func DoTestGetOneConversation(sourceID string, sessionType int) {
//	var test TestGetOneConversationCallBack
//	//GetOneConversation(Friend_uid, SingleChatType, test)
//	open_im_sdk.GetOneConversation(sourceID, sessionType, test)
//
//}
func DoTestCreateTextMessage(text string) string {
	operationID := utils.OperationIDGenerator()
	return open_im_sdk.CreateTextMessage(operationID, text)
}

func DoTestCreateTextMessageReliability(mgr *login.LoginMgr, text string) string {
	operationID := utils.OperationIDGenerator()
	return mgr.Conversation().CreateTextMessage(text, operationID)

}

func DoTestCreateImageMessageFromFullPath() string {
	operationID := utils.OperationIDGenerator()
	return open_im_sdk.CreateImageMessageFromFullPath(operationID, "C:\\Users\\Administrator\\Desktop\\rtc.proto")
	//open_im_sdk.SendMessage(&testSendMsg, operationID, s, , "", utils.StructToJsonString(o))
}

func DoTestCreateOtherMessageFromFullPath() string {
	operationID := utils.OperationIDGenerator()
	return open_im_sdk.CreateFileMessageFromFullPath(operationID, "C:\\Users\\Administrator\\Desktop\\2.txt", "2.txt")
	//open_im_sdk.SendMessage(&testSendMsg, operationID, s, , "", utils.StructToJsonString(o))
}

func DoTestCreateVideoMessageFromFullPath() string {
	operationID := utils.OperationIDGenerator()
	return open_im_sdk.CreateVideoMessageFromFullPath(operationID, "C:\\Users\\Administrator\\Desktop\\video_test.mp4", "mp4", 5, "C:\\Users\\Administrator\\Desktop\\shot.jpg")
}

//func DoTestSetConversationDraft() {
//	var test TestSetConversationDraft
//	open_im_sdk.SetConversationDraft("single_c93bc8b171cce7b9d1befb389abfe52f", "hah", test)
//
//}
type TestSetConversationDraft struct {
}

func (t TestSetConversationDraft) OnError(errCode int32, errMsg string) {
	fmt.Printf("SetConversationDraft , OnError %v\n", errMsg)
}

func (t TestSetConversationDraft) OnSuccess(data string) {
	fmt.Printf("SetConversationDraft , OnSuccess %v\n", data)
}

type GetHistoryCallBack struct {
	OperationID string
	Data        string
}

func (g GetHistoryCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "GetHistoryCallBack err", errCode, errMsg)
}

func (g GetHistoryCallBack) OnSuccess(data string) {
	g.Data = data
	log.Info(g.OperationID, "get History success ", data)
}

type SetAppBadgeCallBack struct {
	OperationID string
}

func (g SetAppBadgeCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "SetAppBadgeCallBack err", errCode, errMsg)
}

func (g SetAppBadgeCallBack) OnSuccess(data string) {
	log.Info(g.OperationID, "SetAppBadgeCallBack success ", data)
}

type UpdateFcmTokenCallBack struct {
	OperationID string
}

func (g UpdateFcmTokenCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "UpdateFcmTokenCallBack err", errCode, errMsg)
}

func (g UpdateFcmTokenCallBack) OnSuccess(data string) {
	log.Info(g.OperationID, "UpdateFcmTokenCallBack success ", data)
}

type FindMessageListCallBack struct {
	OperationID string
}

func (g FindMessageListCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "FindMessageListCallBack err", errCode, errMsg)
}

func (g FindMessageListCallBack) OnSuccess(data string) {
	log.Info(g.OperationID, "FindMessageListCallBack success ", data)
}

type SetMessageReactionExtensionsCallBack struct {
	OperationID string
}

func (g SetMessageReactionExtensionsCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "SetMessageReactionExtensionsCallBack err", errCode, errMsg)
}

func (g SetMessageReactionExtensionsCallBack) OnSuccess(data string) {
	log.Info(g.OperationID, "SetMessageReactionExtensionsCallBack success ", data)
}

type GetHistoryReverseCallBack struct {
	OperationID string
}

func (g GetHistoryReverseCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "GetHistoryReverseCallBack err", errCode, errMsg)
}

func (g GetHistoryReverseCallBack) OnSuccess(data string) {
	log.Info(g.OperationID, "GetHistoryReverseCallBack success ", data)
}

type SearchLocalMessagesCallBack struct {
	OperationID string
}

func (g SearchLocalMessagesCallBack) OnError(errCode int32, errMsg string) {
	log.Info(g.OperationID, "SearchLocalMessagesCallBack err", errCode, errMsg)
}

func (g SearchLocalMessagesCallBack) OnSuccess(data string) {
	fmt.Println(g.OperationID, "SearchLocalMessagesCallBack success ", data)
}

type MsgListenerCallBak struct {
}

func (m *MsgListenerCallBak) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
	//fmt.Println("OnRecvC2CReadReceipt , ", groupMsgReceiptList)
}
func (m *MsgListenerCallBak) OnNewRecvMessageRevoked(messageRevoked string) {
	//fmt.Println("OnNewRecvMessageRevoked , ", messageRevoked)
}

func (m *MsgListenerCallBak) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {

}
func (m *MsgListenerCallBak) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {

}

type BatchMsg struct {
}

func (m *BatchMsg) OnRecvNewMessages(groupMsgReceiptList string) {
	log.Info("OnRecvNewMessages , ", groupMsgReceiptList)
}

func (m *MsgListenerCallBak) OnRecvNewMessage(msg string) {
	var mm sdk_struct.MsgStruct
	err := json.Unmarshal([]byte(msg), &mm)
	if err != nil {
		log.Error("", "Unmarshal failed", err.Error())
	} else {
		//		log.Info("", "recv time: ", time.Now().UnixNano(), "send_time: ", mm.SendTime, " client_msg_id: ", mm.ClientMsgID, "server_msg_id", mm.ServerMsgID)
		RecvMsgMapLock.Lock()
		defer RecvMsgMapLock.Unlock()
		t := SendRecvTime{SendIDRecvID: mm.SendID + mm.RecvID, RecvTime: utils.GetCurrentTimestampByMill()}
		RecvAllMsg[mm.ClientMsgID] = &t
		log.Info("", "OnRecvNewMessage  callback", mm.ClientMsgID, mm.SendID, mm.RecvID)
	}
}

type TestSearchLocalMessages struct {
	OperationID string
}

func (t TestSearchLocalMessages) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "SearchLocalMessages , OnError %v\n", errMsg)
}

func (t TestSearchLocalMessages) OnSuccess(data string) {
	log.Info(t.OperationID, "SearchLocalMessages , OnSuccess %v\n", data)
}

//func DoTestSearchLocalMessages() {
//	var t TestSearchLocalMessages
//	operationID := utils.OperationIDGenerator()
//	t.OperationID = operationID
//	var p sdk_params_callback.SearchLocalMessagesParams
//	//p.SessionType = constant.SingleChatType
//	p.SourceID = "18090680773"
//	p.KeywordList = []string{}
//	p.SearchTimePeriod = 24 * 60 * 60 * 10
//	open_im_sdk.SearchLocalMessages(t, operationID, utils.StructToJsonString(p))
//}

type TestDeleteConversation struct {
	OperationID string
}

func (t TestDeleteConversation) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestDeleteConversation , OnError %v\n", errMsg)
}

func (t TestDeleteConversation) OnSuccess(data string) {
	log.Info(t.OperationID, "TestDeleteConversation , OnSuccess %v\n", data)
}
func DoTestDeleteConversation() {
	var t TestDeleteConversation
	operationID := utils.OperationIDGenerator()
	t.OperationID = operationID
	conversationID := "single_17396220460"
	open_im_sdk.DeleteConversation(t, operationID, conversationID)
}
func (m MsgListenerCallBak) OnRecvC2CReadReceipt(data string) {
	fmt.Println("OnRecvC2CReadReceipt , ", data)
}

func (m MsgListenerCallBak) OnRecvMessageRevoked(msgId string) {
	fmt.Println("OnRecvMessageRevoked ", msgId)
}

type conversationCallBack struct {
	SyncFlag int
}

func (c *conversationCallBack) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {
	panic("implement me")
}

func (c *conversationCallBack) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
	panic("implement me")
}

func (c *conversationCallBack) OnSyncServerProgress(progress int) {
	log.Info("", utils.GetSelfFuncName())
}

func (c *conversationCallBack) OnSyncServerStart() {

}

func (c *conversationCallBack) OnSyncServerFinish() {
	c.SyncFlag = 1
	log.Info("", utils.GetSelfFuncName())

}

func (c *conversationCallBack) OnSyncServerFailed() {
	log.Info("", utils.GetSelfFuncName())
}

func (c *conversationCallBack) OnNewConversation(conversationList string) {
	log.Info("", "OnNewConversation returnList is ", conversationList)
}

func (c *conversationCallBack) OnConversationChanged(conversationList string) {
	log.Info("", "OnConversationChanged returnList is", conversationList)
}

func (c *conversationCallBack) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	log.Info("", "OnTotalUnreadMessageCountChanged returnTotalUnreadCount is ", totalUnreadCount)
}

type testMarkC2CMessageAsRead struct {
}

func (testMarkC2CMessageAsRead) OnSuccess(data string) {
	fmt.Println(" testMarkC2CMessageAsRead  OnSuccess", data)
}

func (testMarkC2CMessageAsRead) OnError(code int32, msg string) {
	fmt.Println("testMarkC2CMessageAsRead, OnError", code, msg)
}

//func DoTestMarkC2CMessageAsRead() {
//	var test testMarkC2CMessageAsRead
//	readid := "2021-06-23 12:25:36-7eefe8fc74afd7c6adae6d0bc76929e90074d5bc-8522589345510912161"
//	var xlist []string
//	xlist = append(xlist, readid)
//	jsonid, _ := json.Marshal(xlist)
//	open_im_sdk.MarkC2CMessageAsRead(test, Friend_uid, string(jsonid))
//}

type SendRecvTime struct {
	SendTime             int64
	SendSeccCallbackTime int64
	RecvTime             int64
	SendIDRecvID         string
}

var SendSuccAllMsg map[string]*SendRecvTime //msgid->send+recv:
var SendFailedAllMsg map[string]string
var RecvAllMsg map[string]*SendRecvTime //msgid->send+recv
var SendMsgMapLock sync.RWMutex
var RecvMsgMapLock sync.RWMutex

func init() {
	SendSuccAllMsg = make(map[string]*SendRecvTime)
	SendFailedAllMsg = make(map[string]string)
	RecvAllMsg = make(map[string]*SendRecvTime)

}

func DoTestSetAppBackgroundStatus(isBackground bool) {
	var testSendMsg TestSendMsgCallBack
	operationID := utils.OperationIDGenerator()
	open_im_sdk.SetAppBackgroundStatus(&testSendMsg, operationID, isBackground)
}

func DoTestSendMsg2(sendId, recvID string) {
	m := "Single chat test" + sendId + ":" + recvID + ":"
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateTextMessage(m)
	log.NewInfo(operationID, "send msg:", s)
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "121313"
	o.Desc = "45464"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))
	log.NewInfo(operationID, utils.GetSelfFuncName(), "success", sendId, recvID)
}

func DoTestSendMsg2Group(sendId, groupID string, index int) {
	m := "test: " + sendId + " : " + groupID + " : " + utils.IntToString(index)
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateTextMessage(m)
	log.NewInfo(operationID, "send msg:", s)
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "Title"
	o.Desc = "Desc"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, "", groupID, utils.StructToJsonString(o))
	log.NewInfo(operationID, utils.GetSelfFuncName(), "success")
}

func DoTestSendMsg2c2c(sendId, recvID string, index int) {
	m := "test: " + sendId + " : " + recvID + " : " + utils.IntToString(index)
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateTextMessage(m)
	log.NewInfo(operationID, "send msg:", s)
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "Title"
	o.Desc = "Desc"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))
	log.NewInfo(operationID, utils.GetSelfFuncName(), "success")
}

type TestMarkGroupMessageAsRead struct {
	OperationID string
}

func (t TestMarkGroupMessageAsRead) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestMarkGroupMessageAsRead , OnError %v\n", errMsg)
}

func (t TestMarkGroupMessageAsRead) OnSuccess(data string) {
	log.Info(t.OperationID, "TestMarkGroupMessageAsRead , OnSuccess %v \n", data)
}
func DoTestMarkGroupMessageAsRead() {
	groupID := "769129589"
	msgIDList := []string{"b4602c5a9f261c3832f6513a21cf55d6"}
	operationID := utils.OperationIDGenerator()
	var testMarkGroupMessageAsRead TestMarkGroupMessageAsRead
	testMarkGroupMessageAsRead.OperationID = operationID
	open_im_sdk.MarkGroupMessageAsRead(&testMarkGroupMessageAsRead, operationID, groupID, utils.StructToJsonString(msgIDList))

}

func DoTestSendMsg(index int, sendId, recvID string, groupID string, idx string) {
	m := "test msg " + sendId + ":" + recvID + ":" + idx
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateTextMessageReliability(allLoginMgr[index].mgr, m)
	var mstruct sdk_struct.MsgStruct
	_ = json.Unmarshal([]byte(s), &mstruct)

	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	testSendMsg.sendTime = utils.GetCurrentTimestampByMill()
	o := server_api_params.OfflinePushInfo{}
	o.Title = "title"
	o.Desc = "desc"
	testSendMsg.sendID = sendId
	testSendMsg.recvID = recvID
	testSendMsg.groupID = groupID
	testSendMsg.msgID = mstruct.ClientMsgID
	log.Info(operationID, "SendMessage", sendId, recvID, groupID, testSendMsg.msgID, index)
	if recvID != "" {
		allLoginMgr[index].mgr.Conversation().SendMessage(&testSendMsg, s, recvID, "", utils.StructToJsonString(o), operationID)
	} else {
		allLoginMgr[index].mgr.Conversation().SendMessage(&testSendMsg, s, "", groupID, utils.StructToJsonString(o), operationID)
	}
	SendMsgMapLock.Lock()
	defer SendMsgMapLock.Unlock()
	x := SendRecvTime{SendTime: utils.GetCurrentTimestampByMill()}
	SendSuccAllMsg[testSendMsg.msgID] = &x
}

//
//func DoTestSendMsgPress(index int, sendId, recvID string, idx string) {
//	m := "test msg " + sendId + ":" + recvID + ":" + idx
//	operationID := utils.OperationIDGenerator()
//	s := DoTestCreateTextMessageReliability(allLoginMgr[index].mgr, m)
//	var mstruct sdk_struct.MsgStruct
//	_ = json.Unmarshal([]byte(s), &mstruct)
//
//	var testSendMsg TestSendMsgCallBackPress
//	testSendMsg.OperationID = operationID
//	o := server_api_params.OfflinePushInfo{}
//	o.Title = "title"
//	o.Desc = "desc"
//	testSendMsg.sendID = sendId
//	testSendMsg.recvID = recvID
//	testSendMsg.msgID = mstruct.ClientMsgID
//
//	log.Warn(operationID, "SendMessage", sendId, recvID, testSendMsg.msgID, index)
//
//	allLoginMgr[index].mgr.Conversation().SendMessage(&testSendMsg, s, recvID, "", utils.StructToJsonString(o), operationID)
//}

func DoTestSendImageMsg(recvID string) {
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateImageMessageFromFullPath()
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "121313"
	o.Desc = "45464"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))
}

func DotestUploadFile() {
	operationID := utils.OperationIDGenerator()
	var testSendMsg TestSendMsgCallBack
	open_im_sdk.UploadFile(&testSendMsg, operationID, "C:\\Users\\Administrator\\Desktop\\video_test.mp4")
}

func DoTestSendOtherMsg(sendId, recvID string) {
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateOtherMessageFromFullPath()
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "121313"
	o.Desc = "45464"
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))
}

func DoTestSendVideo(sendId, recvID string) {
	operationID := utils.OperationIDGenerator()
	s := DoTestCreateVideoMessageFromFullPath()
	var testSendMsg TestSendMsgCallBack
	testSendMsg.OperationID = operationID
	o := server_api_params.OfflinePushInfo{}
	o.Title = "121313"
	o.Desc = "45464"
	log.NewInfo(operationID, s)
	open_im_sdk.SendMessage(&testSendMsg, operationID, s, recvID, "", utils.StructToJsonString(o))

}

type TestClearMsg struct {
	OperationID string
}

func (t *TestClearMsg) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestClearMsg , OnError ", errMsg)
}

func (t *TestClearMsg) OnSuccess(data string) {
	log.Info(t.OperationID, "TestClearMsg , OnSuccess ", data)
}

func DoTestClearMsg() {
	var test TestClearMsg
	operationID := utils.OperationIDGenerator()
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(&test, operationID)

}

type TestModifyGroupMessageReaction struct {
	OperationID string
}

func (t *TestModifyGroupMessageReaction) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, "TestModifyGroupMessageReaction , OnError ", errMsg)
}

func (t *TestModifyGroupMessageReaction) OnSuccess(data string) {
	log.Info(t.OperationID, "TestModifyGroupMessageReaction , OnSuccess ", data)
}

func DoTestGetSelfUserInfo() {
	var test TestModifyGroupMessageReaction
	open_im_sdk.GetSelfUserInfo(&test, "s")
}
