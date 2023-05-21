package db

import (
	"encoding/json"
	"errors"
	"open_im_sdk/pkg/db/model_struct"
	"open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
)

func (d *DataBase) GetMessageReactionExtension(msgID string) (result *model_struct.LocalChatLogReactionExtensions, err error) {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	var l model_struct.LocalChatLogReactionExtensions
	return &l, utils.Wrap(d.conn.Where("client_msg_id = ?",
		msgID).Take(&l).Error, "GetMessageReactionExtension failed")
}

func (d *DataBase) InsertMessageReactionExtension(messageReactionExtension *model_struct.LocalChatLogReactionExtensions) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return utils.Wrap(d.conn.Create(messageReactionExtension).Error, "InsertMessageReactionExtension failed")
}
func (d *DataBase) UpdateMessageReactionExtension(c *model_struct.LocalChatLogReactionExtensions) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	t := d.conn.Updates(c)
	if t.RowsAffected == 0 {
		return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
	}
	return utils.Wrap(t.Error, "UpdateConversation failed")
}
func (d *DataBase) GetAndUpdateMessageReactionExtension(msgID string, m map[string]*server_api_params.KeyValue) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	var temp model_struct.LocalChatLogReactionExtensions
	err := d.conn.Where("client_msg_id = ?",
		msgID).Take(&temp).Error
	if err != nil {
		temp.ClientMsgID = msgID
		temp.LocalReactionExtensions = []byte(utils.StructToJsonString(m))
		return d.conn.Create(&temp).Error
	} else {
		oldKeyValue := make(map[string]*server_api_params.KeyValue)
		_ = json.Unmarshal(temp.LocalReactionExtensions, &oldKeyValue)
		for k, newValue := range m {
			if _, ok := oldKeyValue[k]; ok {
				oldKeyValue[k] = newValue
			}
		}
		temp.LocalReactionExtensions = []byte(utils.StructToJsonString(oldKeyValue))
		t := d.conn.Updates(temp)
		if t.RowsAffected == 0 {
			return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
		}
	}
	return nil
}
func (d *DataBase) DeleteAndUpdateMessageReactionExtension(msgID string, m map[string]*server_api_params.KeyValue) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	var temp model_struct.LocalChatLogReactionExtensions
	err := d.conn.Where("client_msg_id = ?",
		msgID).Take(&temp).Error
	if err != nil {
		return err
	} else {
		oldKeyValue := make(map[string]*server_api_params.KeyValue)
		_ = json.Unmarshal(temp.LocalReactionExtensions, &oldKeyValue)
		for k, _ := range m {
			if _, ok := oldKeyValue[k]; ok {
				delete(oldKeyValue, k)
			}
		}
		temp.LocalReactionExtensions = []byte(utils.StructToJsonString(oldKeyValue))
		t := d.conn.Updates(temp)
		if t.RowsAffected == 0 {
			return utils.Wrap(errors.New("RowsAffected == 0"), "no update")
		}
	}
	return nil
}
