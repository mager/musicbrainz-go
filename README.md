# musicbrainz-go

https://github.com/mager/musicbrainz-go

_A super minimal MusicBrainz library to power beatbrain.xyz_

## Features

- **Recording Search**: Search recordings by ISRC, artist, and track name
- **Bulk Operations**: Efficient bulk ISRC searches
- **Rate Limiting**: Built-in rate limiting and delays
- **Custom User Agents**: Configurable user agent strings
- **Error Handling**: Robust error handling and logging

## Examples

### Basic Recording Search
```bash
cd examples/recording/search_recordings_by_isrc
go run search_recordings_by_isrc.go
```

### Artist and Track Search
```bash
cd examples/recording/search_recordings_by_artist_and_track
go run search_recordings_by_artist_and_track.go
```

### Bulk ISRC Search (New!)
```bash
cd examples/recording/search_recordings_by_bulk_isrc
go run search_recordings_by_bulk_isrc.go
```

### Get Recording Details
```bash
cd examples/recording/get_recording
go run get_recording_with_includes.go
```

## API Methods

- `SearchRecordingsByISRC()` - Search by single ISRC
- `SearchRecordingsByBulkISRC()` - Search by multiple ISRCs efficiently
- `SearchRecordingsByArtistAndTrack()` - Search by artist and track name
- `GetRecording()` - Get detailed recording information

## Development

Ship a new version:

```
make publish
```
