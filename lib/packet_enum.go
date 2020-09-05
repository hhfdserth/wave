//go:generate stringer -type=ServerPacket,ClientPacket -output=packet_enum_string.go
package lib

type IPacket interface {
	Name() string
}
type ServerPacket int64
type ClientPacket int64

// Name returns server packet name
func (sp ServerPacket) Name() string {
	return sp.String()
}

// Name returns client packet name
func (cp ClientPacket) Name() string {
	return cp.String()
}

const (
	OutSaveResource          ClientPacket = 45
	OutSaveNpc               ClientPacket = 31
	OutSaveItem              ClientPacket = 30
	OutSavePuzzle            ClientPacket = 101
	OutSaveSpell             ClientPacket = 33
	OutSaveShop              ClientPacket = 32
	OutSaveAnimation         ClientPacket = 50
	OutProjectilAttack       ClientPacket = 81
	OutSaveProjectil         ClientPacket = 80
	OutSaveMoral             ClientPacket = 82
	OutSaveCard              ClientPacket = 84
	OutSaveCrafting          ClientPacket = 86
	OutSaveMapTopo           ClientPacket = 93
	OutSteamTransaction      ClientPacket = 94
	OutSaveCondition         ClientPacket = 95
	OutSaveConditionVar      ClientPacket = 100
	OutSaveQuest             ClientPacket = 73
	OutCheckAttack           ClientPacket = 12
	OutCheckTopo             ClientPacket = 91
	OutCheckDefensa          ClientPacket = 69
	OutHandlePressEnter      ClientPacket = 14
	OutCastingSpell          ClientPacket = 42
	OutPVP                   ClientPacket = 92
	IncAlertMSG              ServerPacket = 1
	IncLoginOk               ServerPacket = 2
	IncClassesData           ServerPacket = 3
	IncInGame                ServerPacket = 4
	IncPlayerInvUpdate       ServerPacket = 5
	IncPlayerWornEquipment   ServerPacket = 6
	IncPlayerVitals          ServerPacket = 7
	IncCoins                 ServerPacket = 8
	IncSkulls                ServerPacket = 9
	IncPlayerStats           ServerPacket = 10
	IncPlayerData            ServerPacket = 11
	IncPlayerMove            ServerPacket = 12
	IncNpcMove               ServerPacket = 13
	IncPlayerDir             ServerPacket = 14
	IncNpcDir                ServerPacket = 15
	IncPlayerXY              ServerPacket = 16
	IncAttack                ServerPacket = 17
	IncNpcAttack             ServerPacket = 18
	IncCheckMap              ServerPacket = 19
	IncMapData               ServerPacket = 20
	IncMapItemData           ServerPacket = 21
	IncMapNPCData            ServerPacket = 22
	IncSteamAchievement      ServerPacket = 23
	IncGlobalMessage         ServerPacket = 24
	IncPlayerMessage         ServerPacket = 25
	IncMapMessage            ServerPacket = 26
	IncSpawnItem             ServerPacket = 27
	IncUpdateItem            ServerPacket = 28
	IncSpawnNPC              ServerPacket = 29
	IncNpcDead               ServerPacket = 30
	IncUpdateNPC             ServerPacket = 31
	IncUpdateShop            ServerPacket = 32
	IncUpdateSpell           ServerPacket = 33
	IncSpells                ServerPacket = 34
	IncPlayerTimer           ServerPacket = 35
	IncResourceCache         ServerPacket = 36
	IncUpdateResource        ServerPacket = 37
	IncPing                  ServerPacket = 38
	IncActionMessage         ServerPacket = 39
	IncPlayerExp             ServerPacket = 40
	IncPlayerSprite          ServerPacket = 41
	IncUpdateAnimation       ServerPacket = 42
	IncAnimation             ServerPacket = 43
	IncMapNpcVitals          ServerPacket = 44
	IncCooldown              ServerPacket = 45
	IncClearSpellBuffer      ServerPacket = 46
	IncSayMessage            ServerPacket = 47
	IncOpenShop              ServerPacket = 48
	IncInspectPlayer         ServerPacket = 49
	IncStunned               ServerPacket = 50
	IncMapWornEquipment      ServerPacket = 51
	IncBank                  ServerPacket = 52
	IncUpdateNewspaper       ServerPacket = 53
	IncPlayerCurrent         ServerPacket = 54
	IncClose                 ServerPacket = 55
	IncTradeUpdate           ServerPacket = 56
	IncTradeStatus           ServerPacket = 57
	IncGameData              ServerPacket = 58
	IncHighIndex             ServerPacket = 59
	IncTarget                ServerPacket = 60
	IncCharData              ServerPacket = 61
	IncDefensa               ServerPacket = 62
	IncUpdateQuest           ServerPacket = 63
	IncPlayerQuest           ServerPacket = 64
	IncSendTime              ServerPacket = 65
	IncSolarEXP              ServerPacket = 66
	IncPlayerInvite          ServerPacket = 67
	IncPartyUpdate           ServerPacket = 68
	IncPartyVitals           ServerPacket = 69
	IncBossMsg               ServerPacket = 70
	IncPlayerXYMap           ServerPacket = 71
	IncProjectil             ServerPacket = 72
	IncUpdateProjectil       ServerPacket = 73
	IncUpdateMoral           ServerPacket = 74
	IncFlash                 ServerPacket = 75
	IncRequestEditor         ServerPacket = 76
	IncUpdateCard            ServerPacket = 77
	IncPlayerCards           ServerPacket = 78
	IncUpdateCrafting        ServerPacket = 79
	IncPlayerCrafting        ServerPacket = 80
	IncPlayerDeath           ServerPacket = 81
	IncUpdateTitle           ServerPacket = 82
	IncPlayerDataType        ServerPacket = 83
	IncBuffInfo              ServerPacket = 84
	IncSpellAnim             ServerPacket = 85
	IncBattleMsg             ServerPacket = 86
	IncNPCProjectil          ServerPacket = 87
	IncMapSound              ServerPacket = 88
	IncArrow                 ServerPacket = 89
	IncBandit                ServerPacket = 90
	IncPlayerGameData        ServerPacket = 91
	IncPVP                   ServerPacket = 92
	IncMapTopoData           ServerPacket = 93
	IncTopoDrill             ServerPacket = 94
	IncUpdateCondition       ServerPacket = 95
	IncServerMsg             ServerPacket = 96
	IncDailyCheck            ServerPacket = 97
	IncSpawnMobItem          ServerPacket = 98
	IncProjectilCross        ServerPacket = 99
	IncReceiveHour           ServerPacket = 100
	IncReceiveBubble         ServerPacket = 101
	IncChoice                ServerPacket = 102
	IncClearChoice           ServerPacket = 103
	IncGetAdminInfo          ServerPacket = 104
	IncBreakPoint            ServerPacket = 105
	IncUpdateMapEvent        ServerPacket = 106
	IncSendConditionVar      ServerPacket = 107
	IncSendPlayerVars        ServerPacket = 108
	IncQuestMsg              ServerPacket = 109
	IncUpdatePuzzle          ServerPacket = 110
	IncUpdatePuzzleCache     ServerPacket = 111
	IncUpdatePlayerList      ServerPacket = 112
	IncCombo                 ServerPacket = 113
	IncTutorial              ServerPacket = 114
	IncPlayerElement         ServerPacket = 115
	IncPlayerRank            ServerPacket = 116
	IncEarthquake            ServerPacket = 117
	IncEventRank             ServerPacket = 118
	IncTravelInfo            ServerPacket = 119
	IncLoginBackOK           ServerPacket = 120
	IncPlayerAwards          ServerPacket = 121
	IncAwards                ServerPacket = 122
	IncCharacterConfirm      ServerPacket = 123
	IncOpenUI                ServerPacket = 124
	IncLogOutOK              ServerPacket = 125
	IncPlayerScore           ServerPacket = 126
	IncPartyQuestInfo        ServerPacket = 127
	IncDemoUpdateQuest       ServerPacket = 128
	IncCashInvUpdate         ServerPacket = 129
	IncPlayerInfo            ServerPacket = 130
	IncPlayerHonor           ServerPacket = 131
	IncPlayerBuddies         ServerPacket = 132
	IncBuddiesAction         ServerPacket = 133
	IncDrillBreak            ServerPacket = 134
	IncArenaBonus            ServerPacket = 136
	IncNpcBuffInfo           ServerPacket = 137
	IncAnnouncer             ServerPacket = 138
	IncOpenBackground        ServerPacket = 139
	IncCouponReady           ServerPacket = 140
	IncProfessionStart       ServerPacket = 141
	IncModifier              ServerPacket = 142
	IncPlayerMaster          ServerPacket = 143
	IncUrl                   ServerPacket = 144
	IncPlayerStatus          ServerPacket = 145
	IncPlayerConnected       ServerPacket = 146
	IncServerStats           ServerPacket = 147
	IncGetModInfo            ServerPacket = 148
	IncPlayerTrap            ServerPacket = 149
	IncShowEmoticon          ServerPacket = 150
	IncProfessionData        ServerPacket = 151
	IncServerVars            ServerPacket = 152
	IncCashShopData          ServerPacket = 153
	IncBlockList             ServerPacket = 154
	IncConstantData          ServerPacket = 155
	OutMapEditorCancel       ClientPacket = 22
	OutRequestEditor         ClientPacket = 83
	OutSelected              ClientPacket = 74
	OutNewAccount            ClientPacket = 1
	OutAddChar               ClientPacket = 3
	OutLogin                 ClientPacket = 2
	OutGetPing               ClientPacket = 46
	OutChangeSpellSlots      ClientPacket = 70
	OutHotbarChange          ClientPacket = 71
	OutMap                   ClientPacket = 21
	OutPlayerDir             ClientPacket = 10
	OutSayMsg                ClientPacket = 8
	OutReportPlayer          ClientPacket = 26
	OutBan                   ClientPacket = 29
	OutHotkeys               ClientPacket = 13
	OutWarpMeTo              ClientPacket = 15
	OutWarpToMe              ClientPacket = 16
	OutAdminTask             ClientPacket = 17
	OutSetAccess             ClientPacket = 34
	OutTakeCapture           ClientPacket = 43
	OutSpawnItem             ClientPacket = 48
	OutRequestLevelUp        ClientPacket = 51
	OutChangePassword        ClientPacket = 18
	OutLoginBack             ClientPacket = 103
	OutWhisper               ClientPacket = 75
	OutHotbarUse             ClientPacket = 72
	OutUseMegaphone          ClientPacket = 19
	OutClientRevision        ClientPacket = 20
	OutSaveCombo             ClientPacket = 102
	OutPartyLeave            ClientPacket = 79
	OutPartyRequest          ClientPacket = 76
	OutReplyPlayerInvitation ClientPacket = 77
	OutDeclineParty          ClientPacket = 78
	OutRequestData           ClientPacket = 85
	OutEnhancement           ClientPacket = 28
	OutMapRespawn            ClientPacket = 25
	OutTrainStat             ClientPacket = 49
	OutDragItem              ClientPacket = 7
	OutProfessionEnd         ClientPacket = 6
	OutPlayerMsg             ClientPacket = 5
	OutWhosOnline            ClientPacket = 35
	OutTryOpenUI             ClientPacket = 36
	OutBanList               ClientPacket = 27
	OutSwapSlots             ClientPacket = 44
	OutUseItem               ClientPacket = 11
	OutUseCashItem           ClientPacket = 106
	OutDropCalaveras         ClientPacket = 110
	OutDropItem              ClientPacket = 24
	OutBuyItem               ClientPacket = 54
	OutSellItem              ClientPacket = 55
	OutDepositItem           ClientPacket = 57
	OutWithdrawItem          ClientPacket = 58
	OutChangeBankSlots       ClientPacket = 56
	OutCloseBank             ClientPacket = 59
	OutPlayerSearch          ClientPacket = 37
	OutTradeRequest          ClientPacket = 61
	OutAdminWarp             ClientPacket = 60
	OutUnequip               ClientPacket = 47
	OutAcceptTrade           ClientPacket = 62
	OutDeclineTrade          ClientPacket = 63
	OutTradeItem             ClientPacket = 64
	OutUntradeItem           ClientPacket = 65
	OutRequestUseChar        ClientPacket = 66
	OutRequestNewChar        ClientPacket = 67
	OutRequestDelChar        ClientPacket = 68
	OutStopDefence           ClientPacket = 69
	OutCrafting              ClientPacket = 87
	OutTitle                 ClientPacket = 90
	OutDeathType             ClientPacket = 89
	OutChoice                ClientPacket = 96
	OutInteraction           ClientPacket = 97
	OutGetAdminInfo          ClientPacket = 98
	OutBreakPoint            ClientPacket = 99
	OutWarpMap               ClientPacket = 104
	OutLogOut                ClientPacket = 105
	OutGetPlayerInfo         ClientPacket = 107
	OutUpdatePlayerHonor     ClientPacket = 108
	OutUIBuddies             ClientPacket = 109
	OutCoupon                ClientPacket = 111
	OutCashShopUI            ClientPacket = 112
	OutFavAwards             ClientPacket = 113
	OutSelectMinigames       ClientPacket = 114
	OutDemoValue             ClientPacket = 115
	OutForgeAction           ClientPacket = 116
	OutCloseShop             ClientPacket = 53
	OutGetStats              ClientPacket = 38
	OutButton                ClientPacket = 40
	OutPlayerCommand         ClientPacket = 39
	OutCancel                ClientPacket = 117
	OutResetPassword         ClientPacket = 118
	OutGetUsername           ClientPacket = 119
)
