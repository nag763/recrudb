# recrudb

recrudb is a job application management platform that centralizes job offers, scores candidates' profiles against opportunities, and assists in redacting motivation letters. Built with React (TanStack) frontend and Go backend, deployed on Firebase and Google Cloud Run.

## Purpose

The platform was designed to streamline the job application process by:
- **Centralizing job offers** from various sources into a single, searchable database
- **Scoring candidates' profiles** against job requirements using AI-powered analysis to identify the best matches
- **Assisting with motivation letter writing** by generating tailored, professional cover letters based on the candidate's profile and the job offer details

## Project Structure

```
recrudb/
├── front/                  # React frontend (TanStack Router, Vite, Tailwind CSS)
│   ├── src/               # Source code
│   ├── public/            # Static assets
│   └── README.md          # Frontend-specific documentation
├── services/              # Go backend API
│   ├── internal/          # Internal packages
│   │   ├── agent/         # Agent services (CV analyzer, job offers, etc.)
│   │   └── secret/        # Authentication and secret management
│   ├── transport/         # API and MCP transport layers
│   └── main.go            # Entry point
├── terraform/             # Infrastructure as Code
│   ├── api.tf             # Backend API infrastructure
│   ├── firebase.tf        # Firebase configuration
│   ├── github.tf          # GitHub Actions secrets and variables
│   └── ...
└── .github/
    └── workflows/          # CI/CD pipelines
        ├── deploy_back.yml   # Backend deployment
        └── deploy_front.yml  # Frontend deployment
```

## Features

### Core Capabilities
- **CV Analysis & Storage**: Upload and analyze candidate CVs/resumes, extracting skills, experience, and qualifications
- **Job Offer Centralization**: Aggregate and store job offers from multiple sources into a searchable database
- **AI Profile Scoring**: Intelligent matching algorithm that scores candidates against job offers (out of 20) based on skills, company fit, location, and offer type
- **Motivation Letter Generation**: AI-powered cover letter writing that creates tailored, professional documents matching the candidate's profile to specific job offers

### AI Agents
The platform uses specialized AI agents for different tasks:
- **CV Analyzer Agent**: Parses and stores candidate profiles and CVs
- **Job Offer Agent**: Analyzes job offers and calculates compatibility scores with candidate profiles
- **Reachout Agent**: Generates personalized motivation letters based on candidate profile and job offer details

### Technical Features
- **Frontend**: Modern React application with TanStack Router, Vite, and Tailwind CSS
- **Backend**: Go-based API with Chi router and Firestore integration
- **Database**: Google Firestore for storing job offers, candidate profiles, and generated documents
- **Authentication**: Firebase Authentication with Google Cloud integration
- **Infrastructure**: Terraform-managed deployment on Firebase Hosting and Google Cloud Run

## Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              USER INTERFACE                                   │
│                          (React + TanStack Router)                          │
└─────────────────────┬───────────────────────────────────────────┬─────────┘
                      │                                           │
                      ▼                                           ▼
┌─────────────────────────────────────┐     ┌──────────────────────────────────┐
│           Firebase Authentication     │     │         Firebase Hosting          │
│      (User Login, Session Management) │     │      (Static Assets, SPA)          │
└─────────────────────────────────────┘     └──────────────────────────────────┘
                      │                                           │
                      │ User Requests (API Calls)               │
                      ▼                                           ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                         GO BACKEND API (Cloud Run)                            │
│                    ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐         │
│                    │    Router   │  │   Handlers  │  │   MCP Server    │         │
│                    │    (Chi)     │  │  (HTTP/JSON) │  │  (Model Context) │         │
│                    └──────┬──────┘  └──────┬──────┘  └─────────┬────────┘         │
└─────────────────────────┼─────────────────────────┼──────────────────┘
                      │                           │                      │
                      ▼                           ▼                      ▼
            ┌─────────────────────┐    ┌─────────────────────┐    ┌─────────────────┐
            │   CV Analyzer Agent  │    │   Job Offer Agent   │    │  Reachout Agent  │
            │  - Parse CVs         │    │  - Analyze Offers    │    │ - Generate       │
            │  - Extract Skills     │    │  - Score Matches     │    │   Cover Letters  │
            │  - Store Profiles     │    │  - (0-20 scoring)    │    │ - Tailored Docs   │
            └─────────┬────────────┘    └─────────┬────────────┘    └──────┬─────────┘
                      │                         │                      │
                      └─────────────────────────┼──────────────────────┘
                                            │
                                            ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        Google Cloud Firestore                                │
│                  (Job Offers, Candidate Profiles, Documents)                   │
└─────────────────────────────────────────────────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        Google Vertex AI (Gemini Models)                        │
│                  (AI-powered Analysis, Scoring, Generation)                     │
└─────────────────────────────────────────────────────────────────────────────┘
```

## Tech Stack

### Frontend
- **Framework**: React with TanStack Router
- **Build Tool**: Vite
- **Styling**: Tailwind CSS
- **State Management**: TanStack Query, TanStack Store
- **Testing**: Vitest
- **Linting & Formatting**: ESLint, Prettier

### Backend
- **Language**: Go 1.25+
- **Framework**: Chi Router
- **Database**: Google Firestore
- **AI/ML**: Google Vertex AI (Gemini models)
- **Authentication**: Firebase Auth
- **MCP**: Model Context Protocol integration

### Infrastructure
- **Hosting**: Firebase Hosting
- **Backend**: Google Cloud Run
- **IaC**: Terraform
- **CI/CD**: GitHub Actions

## Quick Start

### Prerequisites

- Node.js 20+
- Go 1.25+
- Firebase CLI
- Google Cloud SDK
- Terraform

### Frontend

```bash
cd front
npm install
npm run dev
```

The development server runs on port 3000.

### Backend

```bash
cd services
# Set required environment variables
export PROJECT_NAME=recrudb
export MODEL=gemini-2.0-flash
# Build and run
go build -o server && ./server
```

## Environment Variables

### Frontend (Vite)

Create a `.env` file in the `front/` directory:

```
VITE_API_KEY=<Firebase API Key>
VITE_API_BASE_URL=/api/v1
VITE_PROJECT_NAME=recrudb
```

### Backend (Go)

```
PROJECT_NAME=recrudb
MODEL=gemini-2.0-flash
GOOGLE_APPLICATION_CREDENTIALS=path/to/service-account.json
```

## Deployment

### Frontend to Firebase

```bash
cd front
npm run build
firebase deploy --only hosting
```

### Backend to Cloud Run

The GitHub Actions workflow (`.github/workflows/deploy_back.yml`) automatically deploys the backend to Google Cloud Run on push to main.

### Infrastructure with Terraform

```bash
cd terraform
terraform init
terraform plan
terraform apply
```

## Project Configuration

### Firebase

The project uses Firebase for:
- Hosting (frontend)
- User Authentication
- Firestore database for storing job offers, candidate profiles, and generated documents

Configure in `terraform/firebase.tf` and `front/firebase.json`.

### Google Cloud

Backend services are deployed to Google Cloud Run with Workload Identity Federation for secure CI/CD deployments. The platform uses:
- Vertex AI for AI-powered profile scoring and document generation
- Cloud Firestore for persistent data storage

Configure in `terraform/api.tf` and `terraform/wif.tf`.

## GitHub Actions Secrets

The deployment workflows require the following secrets in GitHub Actions:

- `WIF_PROVIDER`: Workload Identity Federation provider name
- `WIF_SERVICE_ACCOUNT`: Google Cloud service account email
- `FIREBASE_API_KEY`: Firebase API key

These are managed through Terraform in `terraform/github.tf`.

## Security Notes

- The backend uses a dynamically generated internal secret for API authentication
- All sensitive values should be provided through environment variables or secrets management
- The `.env` file pattern is in `.gitignore` and should never be committed
- Firebase and Google Cloud credentials should use Workload Identity Federation

## License

This project is proprietary. Do not use without permission.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## Contact

For questions or support, please refer to the project documentation or open an issue.
