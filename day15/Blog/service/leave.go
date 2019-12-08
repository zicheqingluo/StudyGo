package service

import (
	"studygo/day15/Blog/dao/db"
	"studygo/day15/Blog/model"
)

func GetLeaveList()(leaveList []*model.Leave,err error){
	leaveList,err = db.GetLeaveList()
	if err != nil {
		return
	}
	return
}

func AddLeave(leaveInfo *model.Leave) error{
	lastId,err := db.AddLeave(leaveInfo)
	if err !=nil &&  lastId > 0{
		return nil
	}
	return err


}