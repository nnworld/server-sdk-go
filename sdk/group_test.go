package sdk

import (
	"testing"
)

func TestRongCloud_GroupCreate(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
	)

	err := rc.GroupCreate(
		"u01",
		"rongcloud_group01",
		[]string{"u01", "u02"},
	)

	t.Log(err)
}

func TestRongCloud_GroupGet(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
	)

	rep, err := rc.GroupGet(
		"u01",
	)
	if err == nil {
		t.Log(rep)
	}
	t.Log(err)
}

func TestRongCloud_GroupJoin(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	err := rc.GroupJoin(
		"u01",
		"rongcloud_group01",
		"u03",
	)

	t.Log(err)
}

func TestRongCloud_GroupUpdate(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	err := rc.GroupUpdate(
		"u01",
		"rongcloud_group02",
	)

	t.Log(err)
}

func TestRongCloud_GroupQuit(t *testing.T) {

	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	err := rc.GroupQuit(
		"u03",
		"u01",
	)

	t.Log(err)
}

func TestRongCloud_GroupSync(t *testing.T) {

	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	group := Group{ID: "u02", Name: "rongcloud_group02"}
	groups := []Group{}
	groups = append(groups, group)
	err := rc.GroupSync(
		"u04",
		groups,
	)

	t.Log(err)
}

func TestRongCloud_GroupGagAdd(t *testing.T) {

	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	err := rc.GroupGagAdd(
		"u01",
		[]string{"u02"},
		300,
	)
	t.Log(err)
}

func TestRongCloud_GROUPGagList(t *testing.T) {

	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	rep, err := rc.GroupGagList(
		"u01",
	)
	if err == nil {
		t.Log(rep)
	}
	t.Log(err)
}

func TestRongCloud_GroupGagremove(t *testing.T) {

	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	err := rc.GroupGagRemove(
		"u01",
		[]string{"u02"},
	)

	t.Log(err)
}

func TestRongCloud_GroupDismiss(t *testing.T) {

	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)

	err := rc.GroupDismiss(
		"u01",
		"u01",
	)

	t.Log(err)
}

func TestRongCloud_GroupMuteAllMembersAdd(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	err := rc.GroupMuteAllMembersAdd(
		[]string{
			"group01",
			"group02",
		})
	t.Log(err)
}

func TestRongCloud_GroupMuteAllMembersList(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	group, err := rc.GroupMuteAllMembersGetList(
		[]string{
			"group01",
			"group02",
		})
	if err == nil {
		t.Log(group)
	}
	t.Log(err)
}

func TestRongCloud_GroupMuteAllMembersRemove(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	err := rc.GroupMuteAllMembersRemove(
		[]string{
			"group01",
			"group02",
		})
	t.Log(err)
}

func TestRongCloud_GroupGMuteMembersAdd(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	err := rc.GroupMuteMembersAdd(
		"gourp01",
		[]string{
			"u01",
			"u02",
		},
		30,
	)
	t.Log(err)
}

func TestRongCloud_GroupMuteMembersGetList(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	rep, err := rc.GroupMuteMembersGetList(
		"gourp01",
	)
	if err == nil {
		t.Log(rep)
	}
	t.Log(err)
}

func TestRongCloud_GroupMuteMembersRemove(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	err := rc.GroupMuteMembersRemove(
		"gourp01",
		[]string{
			"u01",
			"u02",
		},
	)
	t.Log(err)
}

func TestRongCloud_GroupMuteWhiteListUserAdd(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	err := rc.GroupMuteWhiteListUserAdd(
		"gourp01",
		[]string{
			"u01",
			"u02",
		},
	)
	t.Log(err)
}

func TestRongCloud_GroupMuteWhiteListUserGetList(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	rep, err := rc.GroupMuteWhiteListUserGetList(
		"gourp01",
	)
	if err == nil {
		t.Log(rep)
	}
	t.Log(err)
}

func TestRongCloud_GroupMuteWhiteListUserRemove(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	err := rc.GroupMuteWhiteListUserRemove(
		"gourp01",
		[]string{
			"u01",
			"u02",
		},
	)
	t.Log(err)
}