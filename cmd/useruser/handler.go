package main

import (
	"context"

	useruser "github.com/808-not-found/tik_duck/kitex_gen/useruser"
)

// UserUserServiceImpl implements the last service interface defined in the IDL.
type UserUserServiceImpl struct{}

// 登录用户对其他用户进行关注或取消关注。
// UserRelationAction implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationAction(
	ctx context.Context,
	req *useruser.RelationActionRequest,
) (resp *useruser.RelationActionResponse, err error) {
	// TODO: Your code here...
	// Request
	// 1: string Token //用户鉴权token
	// 2: i64 ToUserId //对方用户id
	// 3: i32 ActionType // 1-关注，2-取消关注
	// Response
	// 1: i32 StatusCode //状态码，0- 成功，其他值-失败
	// 2: optional string StatusMsg //返回状态描述

	return
}

// 登录用户关注的所有用户列表。
// UserRelationFollowList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFollowList(
	ctx context.Context,
	req *useruser.RelationFollowListRequest,
) (resp *useruser.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	// Request
	// 1: i64 UserId //用户id
	// 2: string Token //用户鉴权token
	// Response
	// 1: i32 StatusCode //状态码，0成功，其他值-失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: list<User> UserList //用户信息列表
	return
}

// 所有关注登录用户的粉丝列表。
// UserRelationFollowerList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFollowerList(
	ctx context.Context,
	req *useruser.RelationFollowerListRequest,
) (resp *useruser.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	// Request
	// 1: i64 UserId //用户id
	// 2: string Token //用户鉴权token
	// Response
	// 1: i32 StatusCode //状态码，0- 成功，其他值失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: list<User> UserList //用户列表
	return
}

// 所有关注登录用户的好友列表。
// UserRelationFriendList implements the UserUserServiceImpl interface.
func (s *UserUserServiceImpl) UserRelationFriendList(
	ctx context.Context,
	req *useruser.RelationFriendListRequest,
) (resp *useruser.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	// Request
	// 1: i64 UserId //用户id
	// 2: string Token //用户鉴权token
	// Response
	// 1: i32 StatusCode //状态码，0- 成功，其他值失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: list<User> UserList //用户列表
	return
}