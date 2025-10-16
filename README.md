# 🌙 Bedtime Calculator

A mini web app that calculates the average bedtime from multiple time entries.

## Quick Start

```bash
# Run the server
go run main.go

# Or use make
make run

# Open in browser
open http://localhost:8080
```

The server will start on port 8080 by default. You can change the port:

```bash
PORT=3000 go run main.go
```

## Features

- **Textarea Input**: Enter multiple bedtimes, one per line in HH:MM format (e.g., 23:00)
- **Real-time Calculation**: The average bedtime is calculated automatically as you type
- **URL Query Parameters**: Pre-fill bedtimes using URL parameters
- **Midnight Wraparound Handling**: Correctly handles bedtimes that span midnight (e.g., 23:00 and 01:00)
- **Input Validation**: Shows error messages for invalid time formats
## Usage

### Basic Usage

1. Open the app in your browser
2. Enter bedtimes in the textarea, one per line:
   ```
   23:00
   22:15
   21:00
   ```
3. The average bedtime is calculated automatically and displayed below

### Using URL Query Parameters

You can pre-fill the textarea with bedtimes using the `bedtimes` query parameter:

```
/bedtime-calculator/?bedtimes=23:00,22:15,21:00
```

**Examples:**

- Evening bedtimes: `?bedtimes=23:00,22:15,21:00`
- Late night mix: `?bedtimes=00:30,01:00,23:45,00:15`
- Single time: `?bedtimes=22:00`

## How It Works

### Average Calculation Algorithm

The app uses a smart algorithm to handle bedtimes that cross midnight:

1. **Parse Times**: Convert HH:MM format to minutes from midnight
2. **Adjust for Bedtime**: Times between 00:00-12:00 are treated as "next day" by adding 24 hours
   - This ensures 23:00 and 01:00 are treated as close together (2 hours apart)
   - Without this, they would be 22 hours apart
3. **Calculate Average**: Sum all adjusted times and divide by count
4. **Convert Back**: Transform the result back to HH:MM format in 24-hour time

### Example Calculation

Input: `23:00, 00:30, 01:00`

1. Convert to minutes: `1380, 30, 60`
2. Adjust (add 1440 to times < 720): `1380, 1470, 1500`
3. Average: `(1380 + 1470 + 1500) / 3 = 1450 minutes`
4. Convert back: `1450 % 1440 = 10 minutes = 00:10`

Result: **00:10** ✓

## Technical Details

### Input Format

- **Valid Format**: `HH:MM` where HH is 00-23 and MM is 00-59
- **Examples**: `23:00`, `01:30`, `12:45`, `00:00`
- **Invalid**: `25:00`, `23:60`, `9:00` (should be `09:00`)


## 🚀 Deployment

### Deploy to Render (Free)

This app is configured for easy deployment to [Render](https://render.com):

1. **Push your code to GitHub**
   ```bash
   git add .
   git commit -m "feat: add Render deployment config"
   git push origin main
   ```

2. **Sign up at [render.com](https://render.com)** (free, no credit card required)

3. **Create a new Web Service**
   - Click "New +" → "Web Service"
   - Connect your GitHub repository
   - Render will auto-detect the `render.yaml` file

4. **Deploy!**
   - Click "Create Web Service"
   - Your app will be live at `https://your-app-name.onrender.com`

**Note:** Free tier apps sleep after 15 minutes of inactivity. First request after sleep takes ~30 seconds.

### Alternative: Manual Docker Deployment

```bash
# Build the Docker image
docker build -t bedtime-calculator .

# Run locally
docker run -p 8080:8080 bedtime-calculator

# Open http://localhost:8080
```

## Running the Server

### Using Go

This app can be served using the included Go server:

```bash
# Run the server
go run main.go

# Or build and run
go build -o bedtime-calculator
./bedtime-calculator
```

The server will start on `http://localhost:8080` by default.

**Custom Port:**

```bash
# Set custom port via environment variable
PORT=3000 go run main.go
```

**Endpoints:**
- `/` - Serves the bedtime calculator app
- `/health` - Health check endpoint (returns OK)

**Using Make:**

```bash
make run      # Run the server
make build    # Build binary
make clean    # Remove binaries
make test     # Test endpoints
make help     # Show all commands
```

### Using Python (Alternative)

For quick testing without Go:

```bash
# Python 3
python3 -m http.server 8080

# Open http://localhost:8080
```

## Development

### File Structure

```
bedtime-calculator/
├── index.html          # Main app file (includes CSS and JS)
├── main.go            # Go web server
├── go.mod             # Go module definition
├── Dockerfile         # Docker configuration
├── render.yaml        # Render deployment config
├── .dockerignore      # Docker ignore file
├── .gitignore         # Git ignore file
├── Makefile           # Build commands
└── README.md          # This file
```

### Testing

```bash
# Start Go server
go run main.go

# Open in browser
# http://localhost:8080/

# Test with query params
# http://localhost:8080/?bedtimes=23:00,22:15,21:00
```

## Browser Support

- Modern browsers (Chrome, Firefox, Safari, Edge)
- Mobile responsive design
- Works standalone or embedded in iframe


