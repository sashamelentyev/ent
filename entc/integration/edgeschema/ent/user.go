// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// Relatives holds the value of the relatives edge.
	Relatives []*User `json:"relatives,omitempty"`
	// LikedTweets holds the value of the liked_tweets edge.
	LikedTweets []*Tweet `json:"liked_tweets,omitempty"`
	// Tweets holds the value of the tweets edge.
	Tweets []*Tweet `json:"tweets,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// JoinedGroups holds the value of the joined_groups edge.
	JoinedGroups []*UserGroup `json:"joined_groups,omitempty"`
	// Friendships holds the value of the friendships edge.
	Friendships []*Friendship `json:"friendships,omitempty"`
	// Relationship holds the value of the relationship edge.
	Relationship []*Relationship `json:"relationship,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*TweetLike `json:"likes,omitempty"`
	// UserTweets holds the value of the user_tweets edge.
	UserTweets []*UserTweet `json:"user_tweets,omitempty"`
	// RolesUsers holds the value of the roles_users edge.
	RolesUsers []*RoleUser `json:"roles_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [12]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// RelativesOrErr returns the Relatives value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RelativesOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Relatives, nil
	}
	return nil, &NotLoadedError{edge: "relatives"}
}

// LikedTweetsOrErr returns the LikedTweets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikedTweetsOrErr() ([]*Tweet, error) {
	if e.loadedTypes[3] {
		return e.LikedTweets, nil
	}
	return nil, &NotLoadedError{edge: "liked_tweets"}
}

// TweetsOrErr returns the Tweets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TweetsOrErr() ([]*Tweet, error) {
	if e.loadedTypes[4] {
		return e.Tweets, nil
	}
	return nil, &NotLoadedError{edge: "tweets"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[5] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// JoinedGroupsOrErr returns the JoinedGroups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) JoinedGroupsOrErr() ([]*UserGroup, error) {
	if e.loadedTypes[6] {
		return e.JoinedGroups, nil
	}
	return nil, &NotLoadedError{edge: "joined_groups"}
}

// FriendshipsOrErr returns the Friendships value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendshipsOrErr() ([]*Friendship, error) {
	if e.loadedTypes[7] {
		return e.Friendships, nil
	}
	return nil, &NotLoadedError{edge: "friendships"}
}

// RelationshipOrErr returns the Relationship value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RelationshipOrErr() ([]*Relationship, error) {
	if e.loadedTypes[8] {
		return e.Relationship, nil
	}
	return nil, &NotLoadedError{edge: "relationship"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikesOrErr() ([]*TweetLike, error) {
	if e.loadedTypes[9] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// UserTweetsOrErr returns the UserTweets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserTweetsOrErr() ([]*UserTweet, error) {
	if e.loadedTypes[10] {
		return e.UserTweets, nil
	}
	return nil, &NotLoadedError{edge: "user_tweets"}
}

// RolesUsersOrErr returns the RolesUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesUsersOrErr() ([]*RoleUser, error) {
	if e.loadedTypes[11] {
		return e.RolesUsers, nil
	}
	return nil, &NotLoadedError{edge: "roles_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		}
	}
	return nil
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return (&UserClient{config: u.config}).QueryGroups(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{config: u.config}).QueryFriends(u)
}

// QueryRelatives queries the "relatives" edge of the User entity.
func (u *User) QueryRelatives() *UserQuery {
	return (&UserClient{config: u.config}).QueryRelatives(u)
}

// QueryLikedTweets queries the "liked_tweets" edge of the User entity.
func (u *User) QueryLikedTweets() *TweetQuery {
	return (&UserClient{config: u.config}).QueryLikedTweets(u)
}

// QueryTweets queries the "tweets" edge of the User entity.
func (u *User) QueryTweets() *TweetQuery {
	return (&UserClient{config: u.config}).QueryTweets(u)
}

// QueryRoles queries the "roles" edge of the User entity.
func (u *User) QueryRoles() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRoles(u)
}

// QueryJoinedGroups queries the "joined_groups" edge of the User entity.
func (u *User) QueryJoinedGroups() *UserGroupQuery {
	return (&UserClient{config: u.config}).QueryJoinedGroups(u)
}

// QueryFriendships queries the "friendships" edge of the User entity.
func (u *User) QueryFriendships() *FriendshipQuery {
	return (&UserClient{config: u.config}).QueryFriendships(u)
}

// QueryRelationship queries the "relationship" edge of the User entity.
func (u *User) QueryRelationship() *RelationshipQuery {
	return (&UserClient{config: u.config}).QueryRelationship(u)
}

// QueryLikes queries the "likes" edge of the User entity.
func (u *User) QueryLikes() *TweetLikeQuery {
	return (&UserClient{config: u.config}).QueryLikes(u)
}

// QueryUserTweets queries the "user_tweets" edge of the User entity.
func (u *User) QueryUserTweets() *UserTweetQuery {
	return (&UserClient{config: u.config}).QueryUserTweets(u)
}

// QueryRolesUsers queries the "roles_users" edge of the User entity.
func (u *User) QueryRolesUsers() *RoleUserQuery {
	return (&UserClient{config: u.config}).QueryRolesUsers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.Grow(14)
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
