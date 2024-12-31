package models

type SpeedyLogFile struct {
	//IdSpeedyLogFile int64           `db:"idSpeedyLogFile" json:"idSpeedyLogFile"`
	Version  *int64           `db:"Version" json:"version"`
	Time     *string          `db:"Time" json:"time"`
	Sessions *[]SpeedySession `db:"Sessions" json:"sessions"`
}

type SpeedySession struct {
	//IdSpeedySession   int64                 `db:"idSpeedySession" json:"idSpeedySession"`
	IsReplay          *bool                  `db:"IsReplay" json:"is_replay"`
	GameModeType      *string                `db:"GameModeType" json:"game_mode"`
	Level             *string                `db:"Level" json:"level"`
	LevelUGCContentID *string                `db:"LevelUGCContentID" json:"level_ugc_content_id"`
	StartTime         *string                `db:"StartTime" json:"start"`
	EndTime           *string                `db:"EndTime" json:"end"`
	Players           *[]SpeedyPlayer        `json:"players"`
	Events            *[]SpeedySessionEvents `json:"events"`
}

type SpeedySessionEvents struct {
	//IdSpeedySessionEvent int64  `db:"idSpeedySessionEvent" json:"idSpeedySessionEvent"`
	Type      *string `db:"SessionEventType" json:"type"`
	Reference *int64  `db:"PlayerReference" json:"reference"`
	ATimeInMS *int64  `db:"aTimeInMS" json:"atime_ms"`
	RTimeInMS *int64  `db:"rTimeInMS" json:"rtime_ms"`
}

type SpeedyPlayer struct {
	//IdPlayer                    int64   `db:"idPlayer" json:"idPlayer"` // speedyplayers
	Reference                   *int64   `db:"PlayerReference" json:"reference"` // speedyplayers
	SteamID                     *string  `db:"SteamID" json:"steamid"`           // speedyplayers
	Index                       *int64   `db:"PlayerArrayIndex" json:"index"`    // playeringamestats
	Name                        *string  `db:"Name" json:"name"`                 // speedyplayers
	IsHosting                   *bool    `db:"IsHosting" json:"is_hosting"`      //playeringamestats
	IsBot                       *bool    `db:"IsBot" json:"is_bot"`              //speedyplayers
	TotalJumps                  *int64   `db:"TotalJumps" json:"jumps"`          //playeringamestats
	TotalHooks                  *int64   `db:"TotalHooks" json:"hooks_shot"`     //playeringamestats
	TotalHit                    *int64   `db:"HooksHit" json:"hooks_hit"`        //playeringamestats
	Overtakes                   *int64   `db:"Overtakes" json:"overtakes"`       //playeringamestats
	DistanceTraveled            *float64 `db:"DistanceTravelled" json:"distance_traveled"`
	SpeedAvg                    *float64 `db:"SpeedAvg" json:"speed_avg"`                                             // playeringamestats
	WeaponCratesDeployed        *int64   `db:"WeaponCratesDeployed" json:"weapon_crate_dropped"`                      // playeringamestats
	WeaponBombsPlanted          *int64   `db:"WeaponCratesDeployed" json:"weapon_bomb_planted"`                       // playeringamestats
	TotalWeaponCrateHits        *int64   `db:"WeaponCratesDeployed" json:"weapon_crate_victims"`                      // playeringamestats
	Points                      *int64   `db:"Points" json:"points"`                                                  // playeringamestats
	TotalWeaponCrateAirHits     *int64   `db:"WeaponCratesHitInAir" json:"weapon_crate_air_victims"`                  // playeringamestats
	TotalWeaponShockwaveHits    *int64   `db:"WeaponShockwaveHitsInAir" json:"weapon_shockwave_fired"`                // playeringamestats
	TotalWeaponShockwaveBlocked *int64   `db:"TotalWeaponShockwaveBlocked" json:"weapon_shockwave_obstacles_removed"` // playeringamestats
	TotalObsticlesHit           *int64   `db:"TotalWeaponShockwaveBlocked" json:"obstacles_hit"`                      // playeringamestats
	TotalSlideTackles           *int64   `db:"TotalWeaponShockwaveBlocked" json:"slide_tackles"`                      // playeringamestats
}

/*

C:/Users/User/Documents/SavedGames/SpeedRunners/CEngineStorage/AllPlayers/Stats/...
May not be valid JSON - May need a sanitization run before parse.

{
	"version": 1,
	"time": "2024-03-28T18:42:56.309Z",
	"sessions": [
		{
			"is_replay": false,
			"game_mode": "Quick",
			"level": "Void",
			"level_ugc_content_id": "792579982",
			"start": "2024-03-28T18:45:12.675Z",
			"end": "2024-03-28T18:47:14.986Z",
			"players": [
				{
					"reference": 0,
					"steamid": "76561199467853904",
					"index": 0,
					"name": "M18",
					"is_hosting": false,
					"is_bot": false,
					"jumps": 51,
					"hooks_shot": 55,
					"hooks_hit": 53,
					"overtakes": 22,
					"distance_traveled": 1987.78,
					"speed_avg": 20.27181,
				},
				{
					"reference": 1,
					"steamid": "76561198056043934",
					"index": 0,
					"name": "2",
					"is_hosting": false,
					"is_bot": false,
					"jumps": 55,
					"hooks_shot": 49,
					"hooks_hit": 40,
					"overtakes": 21,
					"points": 3,
					"game_losses": 1,
					"distance_traveled": 1977.812,
					"speed_avg": 19.96041,
				},
				{
					"reference": 2,
					"steamid": "76561198042306778",
					"index": 0,
					"name": "SmallFries",
					"is_hosting": false,
					"is_bot": false,
					"overtakes": 6,
					"hooks_shot": 7,
					"jumps": 11,
					"obstacles_hit": 2,
					"hooks_hit": 3,
					"distance_traveled": 473.2674,
					"speed_avg": 10.526,
				}
			],
			"events": [
				{ "type": "eliminated", "reference": 0, "atime_ms": 44481, "rtime_ms": 2559502116 },
				{ "type": "point", "reference": 1, "atime_ms": 45970, "rtime_ms": 2559503606 },
				{ "type": "go", "atime_ms": 51286, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 2, "atime_ms": 87455, "rtime_ms": 36169 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 95662, "rtime_ms": 44376 },
				{ "type": "point", "reference": 1, "atime_ms": 97151, "rtime_ms": 45864 },
				{ "type": "go", "atime_ms": 102477, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 2, "atime_ms": 111270, "rtime_ms": 8792 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 111534, "rtime_ms": 9056 },
				{ "type": "point", "reference": 1, "atime_ms": 113031, "rtime_ms": 10553 }
			]
		},
		{
			"is_replay": false,
			"game_mode": "Custom",
			"level": "Shore Enuff",
			"level_ugc_content_id": "951878944",
			"start": "2024-03-28T18:47:35.377Z",
			"end": "2024-03-28T18:49:51.692Z",
			"players": [
				{
					"reference": 0,
					"steamid": "76561198042306778",
					"index": 0,
					"name": "SmallFries",
					"is_hosting": false,
					"is_bot": false,
					"jumps": 18,
					"hooks_shot": 47,
					"hooks_hit": 20,
					"obstacles_hit": 12,
					"distance_traveled": 1441.248,
					"speed_avg": 18.62419,
				},
				{
					"reference": 1,
					"steamid": "76561198042306778",
					"index": 1,
					"name": "SkullDuggery",
					"is_hosting": false,
					"is_bot": true,
					"jumps": 43,
					"hooks_shot": 28,
					"hooks_hit": 25,
					"weapon_crate_dropped": 3,
					"weapon_crate_victims": 2,
					"points": 3,
					"weapon_crate_air_victims": 1,
					"weapon_shockwave_fired": 1,
					"weapon_shockwave_obstacles_removed": 2,
					"distance_traveled": 1690.176,
					"speed_avg": 20.74011,
				},
				{
					"reference": 2,
					"steamid": "76561198042306778",
					"index": 2,
					"name": "Jailbird",
					"is_hosting": false,
					"is_bot": true,
					"jumps": 47,
					"hooks_shot": 20,
					"hooks_hit": 15,
					"obstacles_hit": 1,
					"weapon_bomb_planted": 1,
					"weapon_crate_dropped": 3,
					"weapon_crate_victims": 2,
					"points": 1,
					"distance_traveled": 1625.853,
					"speed_avg": 18.95243,
				},
				{
					"reference": 3,
					"steamid": "76561198042306778",
					"index": 3,
					"name": "Salem",
					"is_hosting": false,
					"is_bot": true,
					"slide_tackles": 1,
					"jumps": 54,
					"hooks_shot": 18,
					"hooks_hit": 15,
					"points": 1,
					"obstacles_hit": 2,
					"distance_traveled": 1532.063,
					"speed_avg": 18.95697,
				}
			],
			"events": [
				{ "type": "go", "atime_ms": 3817, "rtime_ms": 0 },
				{ "type": "go", "atime_ms": 3825, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 24209, "rtime_ms": 20383 },
				{ "type": "eliminated", "reference": 2, "atime_ms": 27817, "rtime_ms": 23991 },
				{ "type": "eliminated", "reference": 3, "atime_ms": 27921, "rtime_ms": 24095 },
				{ "type": "point", "reference": 1, "atime_ms": 29409, "rtime_ms": 25583 },
				{ "type": "go", "atime_ms": 34728, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 41290, "rtime_ms": 6562 },
				{ "type": "eliminated", "reference": 1, "atime_ms": 49769, "rtime_ms": 15041 },
				{ "type": "eliminated", "reference": 2, "atime_ms": 49777, "rtime_ms": 15049 },
				{ "type": "point", "reference": 3, "atime_ms": 51266, "rtime_ms": 16537 },
				{ "type": "go", "atime_ms": 56579, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 2, "atime_ms": 66282, "rtime_ms": 9702 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 74227, "rtime_ms": 17647 },
				{ "type": "eliminated", "reference": 3, "atime_ms": 75474, "rtime_ms": 18894 },
				{ "type": "point", "reference": 1, "atime_ms": 76961, "rtime_ms": 20381 },
				{ "type": "go", "atime_ms": 82280, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 1, "atime_ms": 87611, "rtime_ms": 5330 },
				{ "type": "eliminated", "reference": 3, "atime_ms": 87611, "rtime_ms": 5330 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 102355, "rtime_ms": 20075 },
				{ "type": "point", "reference": 2, "atime_ms": 103840, "rtime_ms": 21560 },
				{ "type": "go", "atime_ms": 109152, "rtime_ms": 0 },
				{ "type": "eliminated", "reference": 0, "atime_ms": 121875, "rtime_ms": 12722 },
				{ "type": "eliminated", "reference": 2, "atime_ms": 125786, "rtime_ms": 16633 },
				{ "type": "eliminated", "reference": 3, "atime_ms": 126265, "rtime_ms": 17112 },
				{ "type": "point", "reference": 1, "atime_ms": 127752, "rtime_ms": 18600 }
			]
		},
		{
			"is_replay": false,
			"game_mode": "Party",
			"level": "Metro",
			"level_ugc_content_id": "0",
			"start": "2024-03-28T19:12:49.813Z",
			"end": "2024-03-28T19:13:00.214Z",
			"players": [
				{
					"reference": 0,
					"steamid": "76561198042306778",
					"index": 0,
					"name": "SpeedRunner",
					"is_hosting": false,
					"is_bot": true,
					"jumps": 2,
					"distance_traveled": 76.12379,
					"speed_avg": 14.98431,
				},
				{
					"reference": 1,
					"steamid": "76561198042306778",
					"index": 1,
					"name": "Unic",
					"is_hosting": false,
					"is_bot": true,
					"jumps": 2,
					"distance_traveled": 75.97636,
					"speed_avg": 14.95529,
				},
				{
					"reference": 2,
					"steamid": "76561198042306778",
					"index": 2,
					"name": "The Falcon",
					"is_hosting": false,
					"is_bot": true,
					"jumps": 3,
					"distance_traveled": 75.49609,
					"speed_avg": 14.86075,
				},
				{
					"reference": 3,
					"steamid": "76561198042306778",
					"index": 3,
					"name": "Gil",
					"is_hosting": false,
					"is_bot": true,
					"jumps": 3,
					"distance_traveled": 75.39444,
					"speed_avg": 14.84074,
				}
			],
			"events": [
				{ "type": "go", "atime_ms": 3515, "rtime_ms": 0 },
				{ "type": "go", "atime_ms": 3523, "rtime_ms": 0 },
				{ "type": "pause", "atime_ms": 8606, "rtime_ms": 5082 }
			]
		}


*/
