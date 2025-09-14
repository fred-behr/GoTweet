# 🐦 GoTweet – Minimal Twitter Clone (Microservices in Go)

## Overview
GoTweet is a lightweight Twitter clone built with Go, showcasing modern microservice architecture.  
Core features: user accounts, posting tweets, following users, viewing a timeline, and notifications.

---

## 🏗️ Architecture
- **User Service**: Manages users, profiles, authentication.
- **Tweet Service**: Handles creating, deleting, and fetching tweets.
- **Follow Service**: Manages follow/unfollow relationships, generates timelines.
- **Notification Service**: Sends async notifications for follows and new tweets.
- **API Gateway**: Single entrypoint (REST/GraphQL) for clients.

---

## 🔑 Features
1. **User Management**
   - Signup / Login
   - JWT Authentication
   - Profile info (username, bio, avatar)

2. **Tweets**
   - Create tweet
   - Delete tweet
   - Get user tweets

3. **Following**
   - Follow / Unfollow
   - Fetch user’s following/followers
   - Generate timeline feed

4. **Notifications**
   - New follower
   - New tweet from followed user

---

## 🛠️ Tech Stack
- **Language**: Go
- **Database**: PostgreSQL (users, tweets, follows)
- **Cache**: Redis (timeline caching, sessions)
- **Messaging**: NATS or Kafka (notifications & async feed updates)
- **API**: REST (external), gRPC (internal)
- **Containerization**: Docker + Docker Compose
- **Auth**: JWT + password hashing (bcrypt)
- **Observability**: Prometheus + Grafana
- **CI/CD**: GitHub Actions (lint, test, build, deploy)

---

## 📂 Services Breakdown
### 1. User Service
- Endpoints:
  - `POST /signup`
  - `POST /login`
  - `GET /user/:id`
- DB: `users(id, username, email, password_hash, bio, avatar_url)`

### 2. Tweet Service
- Endpoints:
  - `POST /tweet`
  - `DELETE /tweet/:id`
  - `GET /user/:id/tweets`
- DB: `tweets(id, user_id, content, created_at)`

### 3. Follow Service
- Endpoints:
  - `POST /follow/:id`
  - `DELETE /unfollow/:id`
  - `GET /user/:id/followers`
  - `GET /user/:id/following`
  - `GET /timeline`
- DB: `follows(follower_id, followee_id)`

### 4. Notification Service
- Events:
  - `UserFollowed`
  - `NewTweet`
- Delivery:
  - Email / WebSocket / Mocked

### 5. API Gateway
- Routes client requests to correct services.
- Handles auth middleware & request validation.

---

## 🚀 Development Roadmap
1. **Week 1** – User Service + Auth  
2. **Week 2** – Tweet Service + Follow Service  
3. **Week 3** – Timeline logic + Notification Service  
4. **Week 4** – API Gateway, Docker Compose, Monitoring, CI/CD  

---

## ✨ Extensions (Optional)
- Likes & retweets  
- Hashtags & search  
- Media uploads (images/videos)  
- WebSocket live feed updates  
