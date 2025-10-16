# 🚀 Render Deployment Guide

## Quick Start

### Step 1: Push to GitHub

```bash
# Initialize git (if not already done)
git init

# Add all files
git add .

# Commit
git commit -m "feat: add Render deployment configuration"

# Add remote (replace with your repository URL)
git remote add origin https://github.com/YOUR_USERNAME/bedtime-calculator.git

# Push to GitHub
git push -u origin main
```

### Step 2: Deploy on Render

1. **Go to [render.com](https://render.com)** and sign up (free, no credit card)

2. **Click "New +"** in the dashboard

3. **Select "Web Service"**

4. **Connect your GitHub repository**
   - Grant Render access to your GitHub account
   - Select the `bedtime-calculator` repository

5. **Render will auto-detect settings from `render.yaml`**
   - Name: `bedtime-calculator` (or customize)
   - Environment: Docker
   - Plan: Free

6. **Click "Create Web Service"**

7. **Wait for deployment** (usually 2-3 minutes)

8. **Your app is live!** 🎉
   - URL: `https://bedtime-calculator-XXXX.onrender.com`
   - Or your custom URL if configured

## Files Created for Deployment

- **`Dockerfile`** - Containerizes the Go application
- **`render.yaml`** - Render configuration (auto-detected)
- **`.dockerignore`** - Files to exclude from Docker build

## What Happens on Render?

1. Render detects the `render.yaml` file
2. Uses `Dockerfile` to build your app
3. Builds a multi-stage Docker image (build + run)
4. Exposes port 8080
5. Runs health checks on `/health` endpoint
6. Provides a public URL

## Free Tier Limitations

- **Sleeps after 15 minutes** of inactivity
- First request after sleep takes **~30 seconds** to wake up
- **750 hours/month** of runtime
- Perfect for demos and personal projects

## Custom Domain (Optional)

1. Go to your service settings on Render
2. Click "Custom Domain"
3. Add your domain and follow DNS instructions

## Environment Variables

The app uses `PORT=8080` by default (configured in `render.yaml`).

To add more variables:
1. Go to your service on Render
2. Click "Environment"
3. Add key-value pairs

## Monitoring

Render provides:
- **Logs**: Real-time application logs
- **Metrics**: CPU, memory usage
- **Events**: Deployment history

Access from your service dashboard.

## Troubleshooting

### Build fails
- Check Render build logs
- Ensure `go.mod` is committed
- Verify Dockerfile syntax

### App doesn't start
- Check if port 8080 is exposed
- Review application logs
- Test health check endpoint

### Template not found
- Ensure `index.html` is in root directory
- Check Dockerfile COPY commands

## Local Docker Testing (Optional)

```bash
# Build
docker build -t bedtime-calculator .

# Run
docker run -p 8080:8080 bedtime-calculator

# Test
curl http://localhost:8080/health
```

## Support

- [Render Documentation](https://render.com/docs)
- [Render Community](https://community.render.com)

