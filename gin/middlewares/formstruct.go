package middlewares

type AddHistoryForm struct {
	ClanName string `form:"clanName" json:"clanName" binding:"required"`
	Time     string `form:"time" json:"time" binding:"required"`
	Uid      uint   `form:"uid" json:"uid"`
	Round    uint   `form:"round" json:"round" binding:"required"`
	Boss     uint   `form:"boss" json:"boss" binding:"required"`
	Dmg      uint   `form:"dmg" json:"dmg" binding:"required"`
	Flag     uint   `form:"flag" json:"flag"`
}
type GetHistoryPayload struct {
	ClanName string `form:"clanName" json:"clanName" binding:"required"`
	Time     string `form:"time" json:"time" binding:"required"`
}
