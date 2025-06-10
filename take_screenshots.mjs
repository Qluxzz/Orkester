/**
 * Updates screenshots stored in screenshots folder
 */
import { chromium } from "playwright"

const pages = [
  { url: "/artist/1", name: "artist-page.png" },
  { url: "/album/1", name: "album-page.png" },
  { url: "", name: "homepage.png" },
  { url: "/search/a", name: "search.png" },
  { url: "/liked-tracks", name: "liked-tracks.png" },
]

const BASE_URL = "http://localhost"

const amountOfTracks = 1000

// Seed fake date
await fetch(`http://localhost/api/v1/scan/fake?amount=${amountOfTracks}`, {
  method: "PUT",
})

// Like some tracks so we can take a screenshot of liked tracks
for (let i = 0; i < 10; ++i) {
  await fetch(
    `${BASE_URL}/api/v1/track/track-${Math.floor(
      Math.random() * amountOfTracks
    )}/like`,
    { method: "PUT" }
  )
}

const browser = await chromium.launch()
const page = await browser.newPage()

for (const { url, name } of pages) {
  await page.goto(`${BASE_URL}${url}`, { waitUntil: "networkidle" })
  await page.screenshot({ path: `screenshots/${name}` })
}
await browser.close()
