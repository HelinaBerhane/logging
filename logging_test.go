package logging

import (
	"bytes"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func captureLog(t *testing.T, fn func()) string {
	t.Helper()

	var buf bytes.Buffer
	originalWriter := log.Writer()
	originalFlags := log.Flags()
	originalPrefix := log.Prefix()

	log.SetOutput(&buf)
	log.SetFlags(0)
	log.SetPrefix("")

	defer func() {
		log.SetOutput(originalWriter)
		log.SetFlags(originalFlags)
		log.SetPrefix(originalPrefix)
	}()

	fn()

	return buf.String()
}

func TestLogging(t *testing.T) {
	originalLevel := Level
	t.Cleanup(func() {
		Level = originalLevel
	})

	t.Run("Set level too low", func(t *testing.T) {
		err := SetLevel(-1)
		require.Error(t, err)
		require.Equal(t, Level, originalLevel)
	})

	t.Run("Set level too high", func(t *testing.T) {
		err := SetLevel(Error + 1)
		require.Error(t, err)
		require.Equal(t, Level, originalLevel)
	})

	t.Run("Default State", func(t *testing.T) {
		got := captureLog(t, func() {
			LogDebug("debug message")
		})
		assert.Contains(t, got, "debug message")

		got = captureLog(t, func() {
			LogInfo("debug message")
		})
		assert.Contains(t, got, "debug message")

		got = captureLog(t, func() {
			LogWarning("debug message")
		})
		assert.Contains(t, got, "debug message")

		got = captureLog(t, func() {
			LogError(errors.New("debug message"))
		})
		assert.Contains(t, got, "debug message")

		// DO EVENTUALLY test LogFatalError - figure out how to handle it calling os.Exit
	})

	t.Run("Set level to Error", func(t *testing.T) {

		err := SetLevel(Error)
		require.NoError(t, err)
		require.Equal(t, Level, Error)

		got := captureLog(t, func() {
			LogDebug("debug message")
		})
		assert.Empty(t, got)

		got = captureLog(t, func() {
			LogInfo("debug message")
		})
		assert.Empty(t, got)

		got = captureLog(t, func() {
			LogWarning("debug message")
		})
		assert.Empty(t, got)

		got = captureLog(t, func() {
			LogError(errors.New("debug message"))
		})
		assert.Contains(t, got, "debug message")

		// DO EVENTUALLY test LogFatalError - figure out how to handle it calling os.Exit
	})
}
