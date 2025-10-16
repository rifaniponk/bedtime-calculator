# 🌙 Bedtime Calculator

A mini web app that calculates the average bedtime from multiple time entries.

## Features

- **Textarea Input**: Enter multiple bedtimes, one per line in HH:MM format (e.g., 23:00)
- **Real-time Calculation**: The average bedtime is calculated automatically as you type
- **URL Query Parameters**: Pre-fill bedtimes using URL parameters
- **Midnight Wraparound Handling**: Correctly handles bedtimes that span midnight (e.g., 23:00 and 01:00)
- **Input Validation**: Shows error messages for invalid time formats
- **WeGuide Integration**: Full integration with WeGuide iframe communication system
- **Local Storage**: Saves data when used within WeGuide platform

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

### WeGuide Events Published

- `app-loaded`: When the app initializes
- `average-calculated`: When average is calculated
- `input-cleared`: When user clears all input
- `loaded-from-query`: When bedtimes are loaded from URL parameters

### Local Storage

When used within WeGuide platform, the app stores:
- `bedtimes`: Array of valid bedtime strings
- `averageBedtime`: Calculated average time
- `lastUpdated`: Timestamp of last update

## Development

### File Structure

```
apps/bedtime-calculator/
├── index.html          # Main app file (includes CSS and JS)
└── README.md          # This file
```

### Building

```bash
# Build this app only
npm run build bedtime-calculator

# Build all apps
npm run build:all
```

### Testing

```bash
# Start dev server
npm run dev

# Open in browser
# http://localhost:3000/bedtime-calculator/

# Test with query params
# http://localhost:3000/bedtime-calculator/?bedtimes=23:00,22:15,21:00
```

### Test Examples

See `/examples/bedtime-calculator-test.html` for interactive test examples.

## Browser Support

- Modern browsers (Chrome, Firefox, Safari, Edge)
- Mobile responsive design
- Works standalone or embedded in iframe

## Credits

Built with ❤️ for WeGuide Health

