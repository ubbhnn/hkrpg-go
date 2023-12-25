package gdconf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type LevelGroup struct {
	GroupId       uint32
	LoadOnInitial bool           `json:"LoadOnInitial"`
	PropList      []*PropList    `json:"PropList"`    // 实体列表
	MonsterList   []*MonsterList `json:"MonsterList"` // 怪物列表
	NPCList       []*NPCList     `json:"NPCList"`     // NPC列表
	AnchorList    []*AnchorList  `json:"AnchorList"`  // 锚点列表
}
type PropList struct {
	ID                       uint32  `json:"ID"`
	PosX                     float64 `json:"PosX"`
	PosY                     float64 `json:"PosY"`
	PosZ                     float64 `json:"PosZ"`
	RotX                     float64 `json:"RotX"`
	RotY                     float64 `json:"RotY"`
	RotZ                     float64 `json:"RotZ "`
	Name                     string  `json:"Name"`
	PropID                   uint32  `json:"PropID"`
	IsOverrideInitLevelGraph bool    `json:"IsOverrideInitLevelGraph"`
	CampID                   uint32  `json:"CampID"`
	EventID                  uint32  `json:"EventID"`
	MapLayerID               uint32  `json:"MapLayerID"`
	AnchorGroupID            uint32  `json:"AnchorGroupID"`
	AnchorID                 uint32  `json:"AnchorID"`
	MappingInfoID            uint32  `json:"MappingInfoID"`
	ChestClosed              string  `json:"ChestClosed"`
	State                    string  `json:"State"`
}
type AnchorList struct {
	ID         uint32  `json:"ID"`
	PosX       float64 `json:"PosX"`
	PosY       float64 `json:"PosY"`
	PosZ       float64 `json:"PosZ"`
	Name       string  `json:"Name"`
	RotX       float64 `json:"RotX"`
	RotY       float64 `json:"RotY"`
	RotZ       float64 `json:"RotZ "`
	MapLayerID uint32  `json:"MapLayerID"`
}

type MonsterList struct {
	ID           uint32  `json:"ID"`
	PosX         float64 `json:"PosX"`
	PosY         float64 `json:"PosY"`
	PosZ         float64 `json:"PosZ"`
	Name         string  `json:"Name"`
	RotX         float64 `json:"RotX"`
	RotY         float64 `json:"RotY"`
	RotZ         float64 `json:"RotZ "`
	NPCMonsterID uint32  `json:"NPCMonsterID"`
	CampID       uint32  `json:"CampID"`
	EventID      uint32  `json:"EventID"`
}

type NPCList struct {
	ID             uint32   `json:"ID"`
	PosX           float64  `json:"PosX"`
	PosY           float64  `json:"PosY"`
	PosZ           float64  `json:"PosZ"`
	Name           string   `json:"Name"`
	RotX           float64  `json:"RotX"`
	RotY           float64  `json:"RotY"`
	RotZ           float64  `json:"RotZ "`
	NPCID          uint32   `json:"NPCID"`
	DialogueGroups []uint32 `json:"DialogueGroups"`
	MapLayerID     uint32   `json:"MapLayerID"`
	BoardShowList  []uint32 `json:"BoardShowList"`
	RaidID         uint32   `json:"RaidID"`
}

func (g *GameDataConfig) loadGroup() {
	g.GroupMap = make(map[uint32]map[uint32]map[uint32]*LevelGroup)
	playerElementsFilePath := g.configPrefix + "LevelOutput/Group"
	files, err := scanFiles(playerElementsFilePath)
	if err != nil {
		logger.Error("error LevelOutput/Group:", err)
		return
	}

	for _, file := range files {
		levelGroup := new(LevelGroup)
		planeId, floorId, groupId := extractNumbers(filepath.Base(file))

		playerElementsFile, err := os.ReadFile(file)
		if err != nil {
			info := fmt.Sprintf("open file error: %v", err)
			panic(info)
		}

		err = hjson.Unmarshal(playerElementsFile, levelGroup)
		if err != nil {
			info := fmt.Sprintf("parse file error: %v", err)
			panic(info)
		}
		levelGroup.GroupId = groupId

		if g.GroupMap[planeId] == nil {
			g.GroupMap[planeId] = make(map[uint32]map[uint32]*LevelGroup)
		}
		if g.GroupMap[planeId][floorId] == nil {
			g.GroupMap[planeId][floorId] = make(map[uint32]*LevelGroup)
		}

		g.GroupMap[planeId][floorId][groupId] = levelGroup
	}

	logger.Info("load %v Groups", len(g.GroupMap))
}

func GetGroupById(planeId, floorId uint32) map[uint32]*LevelGroup {
	return CONF.GroupMap[planeId][floorId]
}

func GetGroupMap() map[uint32]map[uint32]map[uint32]*LevelGroup {
	return CONF.GroupMap
}

func scanFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func extractNumbers(filename string) (uint32, uint32, uint32) {
	filename = strings.TrimSuffix(filename, ".json")

	parts := strings.Split(filename, "_")
	if len(parts) != 4 {
		return 0, 0, 0
	}

	pValueStr := strings.TrimLeft(parts[1], "P")
	fValueStr := strings.TrimLeft(parts[2], "F")
	gValueStr := strings.TrimLeft(parts[3], "G")

	pValue, _ := strconv.ParseUint(pValueStr, 10, 32)
	fValue, _ := strconv.ParseUint(fValueStr, 10, 32)
	gValue, _ := strconv.ParseUint(gValueStr, 10, 32)

	return uint32(pValue), uint32(fValue), uint32(gValue)
}

func GetStateValue(state string) uint32 {
	stateMap := map[string]uint32{
		"Closed":            0,
		"Open":              1,
		"Locked":            2,
		"BridgeState1":      3,
		"BridgeState2":      4,
		"BridgeState3":      5,
		"BridgeState4":      6,
		"CheckPointDisable": 8,
		"CheckPointEnable":  8,
		"TriggerDisable":    10,
		"TriggerEnable":     10,
		"ChestLocked":       12,
		"ChestClosed":       12,
		"ChestUsed":         13,
		"Elevator1":         14,
		"Elevator2":         15,
		"Elevator3":         16,
		"WaitActive":        17,
		"EventClose":        19,
		"EventOpen":         19,
		"Hidden":            20,
		"TeleportGate0":     21,
		"TeleportGate1":     22,
		"TeleportGate2":     23,
		"TeleportGate3":     24,
		"Destructed":        25,
		"CustomState01":     101,
		"CustomState02":     102,
		"CustomState03":     103,
		"CustomState04":     104,
		"CustomState05":     105,
		"CustomState06":     106,
		"CustomState07":     107,
		"CustomState08":     108,
		"CustomState09":     109,
	}

	value, ok := stateMap[state]
	if !ok {
		return 0
	}

	return value
}
