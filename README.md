# 🎬 StreamMovies

**AI-Powered Movie Streaming Platform built with React, Go (Gin), MongoDB, LangChainGo, and OpenAI**

## About

MagicStream is a full-stack movie streaming platform designed to demonstrate how modern web technologies can be combined to build a scalable, AI-powered application.

The platform integrates a React-based frontend for an interactive user experience, a high-performance backend built with Go and the Gin framework, and an AI recommendation engine powered by LangChainGo and OpenAI to provide personalized movie suggestions.

MongoDB serves as the primary database for storing movie metadata, user preferences, and recommendation-related data.

---

## Features

* 🎥 Movie streaming experience powered by React and React Player
* ⚡ High-performance REST APIs built with Go and Gin
* 🤖 AI-powered movie recommendations using LangChainGo and OpenAI
* 🗄️ Scalable data storage with MongoDB
* 📱 Responsive and user-friendly interface
* 🔄 Modular architecture separating frontend, backend, database, and AI services

---

## Tech Stack

| Layer        | Technology          |
| ------------ | ------------------- |
| Frontend     | React, JavaScript   |
| Backend      | Go, Gin (gin-gonic) |
| Database     | MongoDB             |
| AI / LLM     | LangChainGo, OpenAI |
| Media Player | React Player        |

---

## Architecture

```text
┌─────────────────┐
│   React Client  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Go + Gin APIs  │
└───────┬─────────┘
        │
 ┌──────┴──────┐
 ▼             ▼
MongoDB    AI Recommendation
Database   Service
            │
            ▼
      LangChainGo
           +
         OpenAI
```

---

## Core Functionality

### Movie Streaming

Users can browse and stream movies through an intuitive React-based interface with integrated media playback.

### Backend APIs

The Go backend exposes RESTful APIs for managing movie catalogs, user preferences, and recommendation workflows.

### AI Recommendations

The recommendation service leverages LangChainGo and OpenAI to generate personalized movie suggestions based on user interests and viewing patterns.

### Data Management

MongoDB stores movie metadata, user information, and recommendation-related data while supporting scalability and efficient querying.

---

### Prerequisites

* Go
* Node.js
* MongoDB
* OpenAI API Key


---

## Future Enhancements

* User authentication and authorization
* Watchlist and favorites management
* Collaborative filtering recommendations
* Trending and genre-based recommendations
* Deployment using Docker and Kubernetes
* Real-time analytics dashboard

---

## Author

Built to explore full-stack development, scalable backend services, and AI-powered recommendation systems using modern technologies.
