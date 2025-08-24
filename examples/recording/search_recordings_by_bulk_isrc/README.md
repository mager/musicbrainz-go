# Bulk ISRC Search Example

This example demonstrates how to search for multiple recordings by ISRC in a single request using the MusicBrainz API.

## What it does

The bulk ISRC search allows you to search for multiple recordings at once by providing an array of ISRCs. This is more efficient than making individual requests for each ISRC.

## How it works

1. **Creates a client**: Initializes a MusicBrainz client with a custom user agent
2. **Prepares ISRCs**: Defines an array of ISRCs to search for
3. **Performs bulk search**: Uses the `SearchRecordingsByBulkISRC` method to search for all ISRCs at once
4. **Displays results**: Shows all found recordings with their details

## API Query

The bulk search constructs a query like:
```
isrc:(ISRC1 OR ISRC2 OR ISRC3 OR ISRC4)
```

This is equivalent to searching for any recording that matches any of the provided ISRCs.

## Benefits

- **Efficiency**: Single HTTP request instead of multiple requests
- **Rate limiting**: Respects MusicBrainz rate limits better
- **Performance**: Faster than sequential individual searches
- **Batch processing**: Ideal for applications that need to look up many recordings

## Usage

```bash
cd musicbrainz-go/examples/recording/search_recordings_by_bulk_isrc
go run search_recordings_by_bulk_isrc.go
```

## Example Output

```
Searching for recordings with ISRCs: [USRC12345678 USRC87654321 GBUM71402401 USRC12345679]

Search completed!
Total recordings found: 2
ISRCs searched: 4

Found recordings:
================
1. ID: 12345678-1234-1234-1234-123456789012
   Title: Let It Be
   Artist: The Beatles
   Release Date: 1970-05-08
   Genres: Rock, Pop

2. ID: 87654321-4321-4321-4321-210987654321
   Title: Another Song
   Artist: Some Artist
   Release Date: 2020-01-01
   Genres: Pop
```

## Notes

- The current implementation returns all matching recordings
- To map specific ISRCs to recordings, you would need to parse the response and extract ISRC information
- Consider using additional includes in the search to get more detailed information
- Always respect MusicBrainz rate limits and terms of service
