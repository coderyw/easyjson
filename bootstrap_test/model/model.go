// Package model
// @Author: yinwei
// @File: model
// @Version: 1.0.0
// @Date: 2024/11/28 14:11

package model

type Model1 struct {
	Ab   string `json:"ab"`
	Name int    `json:"name,nomarshal"`
}
