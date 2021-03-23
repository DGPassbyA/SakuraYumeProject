package middlewares

type AddHistoryForm struct {
	ClanName string `form:"clanName" json:"clanName" binding:"required"`
	Time     string `form:"time" json:"time" binding:"required"`
	Uid      int    `form:"uid" json:"uid" binding:"required"`
	Round    int    `form:"round" json:"round" binding:"required"`
	Boss     int    `form:"boss" json:"boss" binding:"required"`
	Dmg      int    `form:"dmg" json:"dmg" binding:"required"`
	Flag     int    `form:"flag" json:"flag" binding:"required"`
}
type GetHistoryPayload struct {
	ClanName string `form:"clanName" json:"clanName" binding:"required"`
	Time     string `form:"time" json:"time" binding:"required"`
}
