// This is called BEFORE your Elm app starts up
//
// The value returned here will be passed as flags
// into your `Shared.init` function.

const volume = (() => {
  const v = localStorage.getItem("volume")
  if (v === null) return null

  const v1 = parseInt(v)

  if (isNaN(v1)) return null

  return v1
})()

export const flags = ({ env }: ElmLand.FlagsArgs) => {
  return { volume, repeat: localStorage.getItem("repeat") }
}

// This is called AFTER your Elm app starts up
//
// Here you can work with `app.ports` to send messages
// to your Elm application, or subscribe to incoming
// messages from Elm
export const onReady = ({ app, env }: ElmLand.OnReadyArgs) => {
  const audio = new Audio()
  if (volume) audio.volume = volume / 100

  function play() {
    audio.play().then(() => app.ports?.stateChange?.send?.("play"))
  }

  function pause() {
    audio.pause()
    app.ports?.stateChange?.send?.("pause")
  }

  audio.addEventListener("ended", (_) => {
    app.ports?.stateChange?.send?.("ended")
  })

  audio.addEventListener("timeupdate", (_) => {
    app.ports?.progressUpdated?.send?.(Math.floor(audio.currentTime))
  })

  app.ports?.setVolume?.subscribe?.((volume) => {
    let v = volume as number
    audio.volume = v / 100
    localStorage.setItem("volume", v.toString())
  })

  app.ports?.setRepeatMode?.subscribe?.((mode) => {
    let m = mode as string
    localStorage.setItem("repeat", m)
  })

  app.ports?.playTrack?.subscribe?.((trackId) => {
    audio.src = `/api/v1/track/${trackId}/stream`
    audio
      .play()
      .then(() => {
        app.ports?.stateChange?.send?.("play")
      })
      .catch((error) => {
        if (error.name === "AbortError") return

        // Uncomment this error in development
        // To see player when trying to play fake tracks
        // if (error.name === "NotSupportedError")
        // 	return

        app.ports?.playbackFailed?.send?.(error.message)
      })
  })

  navigator.mediaSession.setActionHandler("play", play)
  navigator.mediaSession.setActionHandler("pause", pause)
  navigator.mediaSession.setActionHandler("previoustrack", () =>
    app.ports?.stateChange?.send?.("previoustrack")
  )
  navigator.mediaSession.setActionHandler("nexttrack", () =>
    app.ports?.stateChange?.send?.("nexttrack")
  )

  app.ports?.seek?.subscribe?.((timestamp) => {
    let t = timestamp as number
    audio.fastSeek(t)
  })

  app.ports?.play?.subscribe?.(play)
  app.ports?.pause?.subscribe?.(pause)
}

// Type definitions for Elm Land
namespace ElmLand {
  export type FlagsArgs = {
    env: Record<string, string>
  }
  export type OnReadyArgs = {
    env: Record<string, string>
    app: { ports?: Record<string, Port | undefined> }
  }
  export type Port = {
    send?: (data: unknown) => void
    subscribe?: (callback: (data: unknown) => unknown) => void
  }
}
