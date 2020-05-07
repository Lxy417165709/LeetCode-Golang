package main

import "time"

const (
	maxCountOfShowingTweetsPerTime = 10
)

// ---------- Tweet ----------
type Tweet struct {
	id               int
	releaseTimeStamp int64
}

func NewTweet(id int) *Tweet {
	return &Tweet{id, time.Now().UnixNano()}
}
func (t *Tweet) IsNewerThan(referenceTimeStamp int64) bool {
	return t.releaseTimeStamp > referenceTimeStamp
}
func (t *Tweet) GetReleaseTimeStamp() int64 {
	return t.releaseTimeStamp
}
func (t *Tweet) GetId() int {
	return t.id
}

// ---------- Node ----------
type Node struct {
	Tweet *Tweet
	Next  *Node
}

func NewNode(tweet *Tweet) *Node {
	return &Node{tweet, nil}
}

// ---------- TweetList ----------
type TweetList struct {
	dummyHead    *Node
	readPosition *Node
}

func NewTweetList() *TweetList {
	return &TweetList{&Node{}, nil}
}
func (tl *TweetList) InsertFront(post *Tweet) {
	newNode := NewNode(post)
	newNode.Next = tl.dummyHead.Next
	tl.dummyHead.Next = newNode
}
func (tl *TweetList) GetReadingTweet() *Tweet {
	return tl.readPosition.Tweet
}
func (tl *TweetList) ReadInit() {
	tl.readPosition = tl.dummyHead.Next
}
func (tl *TweetList) ReadNext() {
	tl.readPosition = tl.readPosition.Next
}
func (tl *TweetList) ReadIsOver() bool {
	return tl.readPosition == nil
}

// ---------- User ----------
type User struct {
	id        int
	idols     map[int]bool // idols[i] is true -> user has followed the user whose uid is i
	tweetList *TweetList
}

func NewUser(id int) *User {
	return &User{id, make(map[int]bool), NewTweetList()}
}
func (u *User) GetTweetList() *TweetList {
	return u.tweetList
}
func (u *User) Follow(userId int) {
	if u.isMyself(userId) {
		return
	}
	u.idols[userId] = true
}
func (u *User) UnFollow(userId int) {
	if u.isMyself(userId) {
		return
	}
	delete(u.idols, userId)
}
func (u *User) PostTweet(tweetId int) {
	u.tweetList.InsertFront(NewTweet(tweetId))
}
func (u *User) isMyself(userId int) bool {
	return userId == u.id
}

// ---------- Twitter ----------
type Twitter struct {
	users map[int]*User // uid -> user
}

func Constructor() Twitter {
	return Twitter{make(map[int]*User)}
}
func (tt *Twitter) GetNewsFeed(userId int) []int {
	if !tt.userIsExist(userId) {
		tt.generateUser(userId)
	}
	tweetLists := tt.getUserAndHisIdolTweetLists(userId)
	tweetIds := tt.getRecentReleasedTweetIds(tweetLists)
	return tweetIds
}
func (tt *Twitter) Follow(followerId int, followeeId int) {
	if !tt.userIsExist(followerId) {
		tt.generateUser(followerId)
	}
	if !tt.userIsExist(followeeId) {
		tt.generateUser(followeeId)
	}
	user := tt.getUser(followerId)
	user.Follow(followeeId)
}
func (tt *Twitter) Unfollow(followerId int, followeeId int) {
	if !tt.userIsExist(followerId) {
		tt.generateUser(followerId)
	}
	if !tt.userIsExist(followeeId) {
		tt.generateUser(followeeId)
	}
	user := tt.getUser(followerId)
	user.UnFollow(followeeId)
}
func (tt *Twitter) PostTweet(userId int, tweetId int) {
	if !tt.userIsExist(userId) {
		tt.generateUser(userId)
	}
	user := tt.getUser(userId)
	user.PostTweet(tweetId)
}

func (tt *Twitter) getUserAndHisIdolTweetLists(userId int) []*TweetList {
	user := tt.getUser(userId)
	tweetLists := make([]*TweetList, 0)
	for idolId := range user.idols {
		idol := tt.getUser(idolId)
		idolTweetList := idol.GetTweetList()
		idolTweetList.ReadInit()
		tweetLists = append(tweetLists, idolTweetList)
	}
	userTweetList := user.GetTweetList()
	userTweetList.ReadInit()
	tweetLists = append(tweetLists, userTweetList)
	return tweetLists
}

// the function is to get recent Tweet's ids
// using the sort algorithm of link lists
func (tt *Twitter) getRecentReleasedTweetIds(tweetLists []*TweetList) []int {
	tweetIds := make([]int, 0)
	for {
		listIndexOfNewestTweet := getListIndexOfNewestTweet(tweetLists)
		if allListsIsReadOver(tweetLists) {
			break
		}
		readingTweet := tweetLists[listIndexOfNewestTweet].GetReadingTweet()
		tweetIds = append(tweetIds, readingTweet.GetId())
		tweetLists[listIndexOfNewestTweet].ReadNext()
		if quantityIsEnough(tweetIds) {
			break
		}
	}
	return tweetIds
}
func (tt *Twitter) getUser(userId int) *User {
	return tt.users[userId]
}
func (tt *Twitter) userIsExist(userId int) bool {
	return tt.users[userId] != nil
}
func (tt *Twitter) generateUser(userId int) {
	tt.users[userId] = NewUser(userId)
}

func getListIndexOfNewestTweet(tweetLists []*TweetList) int {
	listIndexOfNewestTweet := 0
	recentTimeStamp := int64(0)
	for index, tweetList := range tweetLists {
		if tweetList.ReadIsOver() {
			continue
		}
		readingTweet := tweetList.GetReadingTweet()
		if readingTweet.IsNewerThan(recentTimeStamp) {
			listIndexOfNewestTweet = index
			recentTimeStamp = readingTweet.GetReleaseTimeStamp()
		}
	}
	return listIndexOfNewestTweet
}
func allListsIsReadOver(tweetLists []*TweetList) bool {
	for _, tweetList := range tweetLists {
		if !tweetList.ReadIsOver() {
			return false
		}
	}
	return true
}
func quantityIsEnough(tweetIds []int) bool {
	return len(tweetIds) >= maxCountOfShowingTweetsPerTime
}
