import { test, expect } from "@playwright/test"
import search from "./search.json"
import album from "./album.json"
import artist from "./artist.json"
import liked from "./liked-tracks.json"

test("has title", async ({ page }) => {
  await page.goto("http://localhost:1234")

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/Orkester/)
})

test("search", async ({ page }) => {
  await page.route("*/**/api/v1/search/*", async (route) => {
    await route.fulfill({
      json: search,
    })
  })

  await page.goto("http://localhost:1234")

  await page.getByRole("link", { name: "Search" }).click()

  await expect(page.getByRole("search")).toBeVisible()

  await page.getByRole("search").fill("david")

  await expect(
    page.getByRole("main").getByRole("heading", { name: "Tracks" })
  ).toBeVisible()
  await expect(
    page.getByRole("main").getByRole("heading", { name: "Albums" })
  ).toBeVisible()
  await expect(
    page.getByRole("main").getByRole("heading", { name: "Artists" })
  ).toBeVisible()
})

test("album page", async ({ page }) => {
  await page.route("*/**/api/v1/album/*", async (route) => {
    await route.fulfill({
      json: album,
    })
  })

  await page.goto("http://localhost:1234/album/1")

  await expect(
    page
      .getByRole("main")
      .getByRole("heading", { name: "Hurry Up, We're Dreaming" })
  ).toBeVisible()
})

test("artist page", async ({ page }) => {
  await page.route("*/**/api/v1/artist/*", async (route) => {
    await route.fulfill({
      json: artist,
    })
  })

  await page.goto("http://localhost:1234/artist/1")

  await expect(
    page.getByRole("main").getByRole("heading", { name: "Ruelle" })
  ).toBeVisible()
})

test("liked tracks", async ({ page }) => {
  await page.route("*/**/api/v1/playlist/liked", async (route) => {
    await route.fulfill({
      json: liked,
    })
  })

  await page.goto("http://localhost:1234/liked-tracks")

  await expect(
    page.getByRole("main").getByRole("heading", { name: "Liked tracks" })
  ).toBeVisible()
})

test("player bar", async ({ page }) => {
  await page.route("*/**/api/v1/playlist/liked", async (route) => {
    await route.fulfill({
      json: liked,
    })
  })

  await page.goto("http://localhost:1234/liked-tracks")

  await expect(
    page.getByRole("main").getByRole("heading", { name: "Liked tracks" })
  ).toBeVisible()

  await page
    .getByRole("main")
    .locator("tbody")
    .locator("tr")
    .first()
    .locator("td")
    .first()
    .click()

  await expect(
    page.getByRole("status").getByRole("heading", { name: "Oblivion" })
  ).toBeVisible()
})
