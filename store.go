package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/vyuvaraj/ServShared"
)

type UserStore interface {
	LoadUsers() (map[string]User, error)
	SaveUsers(users map[string]User) error
	LoadKeys() (map[string]*APIKey, error)
	SaveKeys(keys map[string]*APIKey) error
	LoadSessions() (map[string]*Session, error)
	SaveSessions(sessions map[string]*Session) error
}

type ServStoreUserStore struct {
	client *ServShared.StoreClient
}

func NewServStoreUserStore(client *ServShared.StoreClient) *ServStoreUserStore {
	return &ServStoreUserStore{client: client}
}

func (s *ServStoreUserStore) LoadUsers() (map[string]User, error) {
	if s.client == nil {
		return nil, errors.New("store client not initialized")
	}
	data, err := s.client.Get("serv-auth-users", "users.json")
	if err != nil {
		return nil, err
	}
	var users map[string]User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	log.Printf("[PERSISTENCE] Loaded %d users from ServStore", len(users))
	return users, nil
}

func (s *ServStoreUserStore) SaveUsers(users map[string]User) error {
	if s.client == nil {
		return errors.New("store client not initialized")
	}
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return s.client.Put("serv-auth-users", "users.json", data)
}

func (s *ServStoreUserStore) LoadKeys() (map[string]*APIKey, error) {
	if s.client == nil {
		return nil, errors.New("store client not initialized")
	}
	data, err := s.client.Get("serv-auth-users", "apikeys.json")
	if err != nil {
		return nil, err
	}
	var keys map[string]*APIKey
	if err := json.Unmarshal(data, &keys); err != nil {
		return nil, err
	}
	log.Printf("[PERSISTENCE] Loaded %d API keys from ServStore", len(keys))
	return keys, nil
}

func (s *ServStoreUserStore) SaveKeys(keys map[string]*APIKey) error {
	if s.client == nil {
		return errors.New("store client not initialized")
	}
	data, err := json.Marshal(keys)
	if err != nil {
		return err
	}
	return s.client.Put("serv-auth-users", "apikeys.json", data)
}

func (s *ServStoreUserStore) LoadSessions() (map[string]*Session, error) {
	if s.client == nil {
		return nil, errors.New("store client not initialized")
	}
	data, err := s.client.Get("serv-auth-users", "sessions.json")
	if err != nil {
		return nil, err
	}
	var sessions map[string]*Session
	if err := json.Unmarshal(data, &sessions); err != nil {
		return nil, err
	}
	log.Printf("[PERSISTENCE] Loaded %d sessions from ServStore", len(sessions))
	return sessions, nil
}

func (s *ServStoreUserStore) SaveSessions(sessions map[string]*Session) error {
	if s.client == nil {
		return errors.New("store client not initialized")
	}
	data, err := json.Marshal(sessions)
	if err != nil {
		return err
	}
	return s.client.Put("serv-auth-users", "sessions.json", data)
}
