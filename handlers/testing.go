package handlers

import (
	"bytes"
	"context"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"orkester/ent"
	"orkester/indexFiles"
	"orkester/repositories"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const fakePath = "fakePath"

var albumNames []string = []string{
	"Mezzanine",
	"Sledgehammer (From The Motion Picture \"Star Trek Beyond\")",
	"Genghis Khan",
	"La Roux",
	"Chapter One",
	"Up From The Skies: The Polydor Years",
	"Covers",
	"Take It All",
	"Up in Flames",
	"Our Glorious Leader (Japanese Trump Commercial Theme)",
	"Thank Your Lucky Stars",
	"Starboy",
	"Analog Is On",
	"Built on Glass",
	"Starboy",
	"Hôtel Costes 14",
	"Oblivion (Original Motion Picture Soundtrack)",
	"Hurry up, We're Dreaming",
	"Moog Indigo",
	"Jeff Lynne's ELO - Alone in the Universe",
	"Can't Fight Fate (Expanded Edition)",
	"saintmotelevision",
	"My Type EP",
	"Miman",
	"Up & Away",
	"Amnesiac",
	"Once Upon a Dream (from \"Maleficent\") [Original Motion Picture Soundtrack]",
	"Born To Die - The Paradise Edition",
	"BADLANDS (Deluxe)",
	"iii",
	"D.A.N.C.E",
	"Civilization",
	"Audio, Video, Disco.",
	"Vild Honung",
	"Imagineering",
	"All That I Am",
	"Shaman",
	"Freedom",
	"Memorial Beach (Deluxe Edition; 2015 Remaster)",
	"East of the Sun, West of the Moon",
	"Notorious (Deluxe Edition)",
	"Seven and the Ragged Tiger (Deluxe Edition)",
	"Lost And Found Volume 1 : Imagination",
	"Visuals",
	"Songs From The Big Chair",
	"Music From Baz Luhrmann's Film The Great Gatsby (International Streaming Version)",
	"Stoney (Deluxe)",
	"UNDERTALE Soundtrack",
	"Sixteen Saltines",
	"Leave a Trace (Goldroom Remix)",
	"Santana IV",
	"Watch Out!",
	"Skifs Hits!",
	"Forever Your Girl",
	"Show You The Way",
	"What's All The Mumble About",
	"Selected Ambient Works 85-92",
	"Friend Zone",
	"UNDERTALE Soundtrack",
	"Junk",
	"Drukqs",
	"Lockjaw",
	"Guardians of the Zone",
	"Mai Lan",
	"Street Fighting Years",
	"American V: A Hundred Highways",
	"Out of the Blue",
	"Supertramp",
	"Indelibly Stamped",
	"A Momentary Lapse of Reason",
	"The Division Bell",
	"Rattle That Lock (Deluxe)",
	"American IV: The Man Comes Around",
	"Fallen Light",
	"Mirror's Edge (Original Videogame Score)",
	"Maggot Brain",
	"Justice",
	"Woman",
	"Let Me Up (I've Had Enough)",
	"You're Gonna Get it",
	"Festival",
	"Musicians for Le Bonheur 2015",
}

var artistNames []string = []string{
	"Massive Attack",
	"Rihanna",
	"Miike Snow",
	"La Roux",
	"Lemaitre",
	"Jennie A.",
	"Ellen McIlwaine",
	"Placebo",
	"Ruelle",
	"Mike Diva",
	"Beach House",
	"The Weeknd",
	"Michael Gray",
	"Chet Faker",
	"Flight Facilities",
	"M83",
	"Susanne Sundfør",
	"Jean-Jacques Perrey",
	"Fatboy Slim",
	"Electric Light Orchestra",
	"Taylor Dayne",
	"Saint Motel",
	"Nicole Sabouné",
	"Can't Stop Won't Stop",
	"June",
	"Radiohead",
	"Lana Del Rey",
	"Halsey",
	"Justice",
	"Björn Skifs",
	"Module",
	"Santana",
	"a-ha",
	"Duran Duran",
	"Re-Flex",
	"Mew",
	"Tears For Fears",
	"Post Malone",
	"Justin Bieber",
	"Toby Fox",
	"Jack White",
	"CHVRCHES",
	"Goldroom",
	"Blue Swede",
	"Paula Abdul",
	"Thundercat",
	"Michael McDonald",
	"Kenny Loggins",
	"Jacle Bow",
	"Aphex Twin",
	"MAI LAN",
	"Flume",
	"TWRP",
	"Ninja Sex Party",
	"Simple Minds",
	"Johnny Cash",
	"Supertramp",
	"Pink Floyd",
	"David Gilmour",
	"Solar Fields",
	"Phaeleh",
	"Soundmouse",
	"Funkadelic",
	"Tom Petty and the Heartbreakers",
	"Forrister",
}

var trackTitles []string = []string{
	"Teardrop",
	"Sledgehammer - From The Motion Picture \"Star Trek Beyond\"",
	"Genghis Khan",
	"In For The Kill",
	"Closer",
	"Can't Find My Way Home",
	"Running Up That Hill",
	"Take It All",
	"Until We Go Down",
	"Our Glorious Leader (Japanese Trump Commercial Theme) - Original Mix",
	"Majorette",
	"Secrets",
	"The Weekend - Radio Edit",
	"Gold",
	"True Colors",
	"Crave You",
	"Oblivion",
	"Reunion",
	"E.V.A.",
	"The Sun Will Shine on You",
	"All My Life",
	"Up All Night",
	"Move",
	"My Type",
	"Right Track",
	"Up & Away",
	"You And Whose Army?",
	"Once Upon a Dream - From \"Maleficent\" / Pop Version",
	"Dark Paradise",
	"Castle",
	"Hurricane",
	"Gasoline",
	"Control",
	"My Trigger",
	"D.A.N.C.E - Radio Edit",
	"Civilization",
	"Audio, Video, Disco.",
	"Stanna",
	"Sunrise Andromeda",
	"Hermes",
	"Adouma",
	"Once It's Gotcha",
	"Lie Down in Darkness - 2015 Remaster",
	"East of the Sun",
	"Notorious - 2010 Remaster",
	"The Reflex - 2010 Remaster",
	"(I'm Looking For) Cracks in the Pavement - 2010 Remaster",
	"The Politics of Dancing",
	"Carry Me to Safety",
	"Shout",
	"Young And Beautiful",
	"Broken Whiskey Glass",
	"Deja Vu",
	"MEGALOVANIA",
	"It's Raining Somewhere Else",
	"Love Is Blindness",
	"Leave a Trace - Goldroom Remix",
	"Yambu",
	"La Booga Rooga",
	"Right Where We Left Off",
	"I Could Never Leave You",
	"Opposites Attract",
	"Show You The Way",
	"High for You Lover",
	"Ageispolis",
	"Friend Zone",
	"Spear of Justice",
	"For the Kids",
	"Solitude",
	"Laser Gun",
	"Avril 14th",
	"Drop The Game",
	"The No Pants Dance",
	"Easy",
	"Belfast Child - Remastered 2002",
	"God's Gonna Cut You Down",
	"Standin' in the Rain",
	"Maybe I'm A Beggar",
	"Shadow Song",
	"Aries",
	"Travelled",
	"Learning to Fly",
	"Coming Back to Life",
	"5 A.M.",
	"Shard",
	"Introduction",
	"Hurt",
	"Afterglow",
	"Edge Flight",
	"Heat",
	"Hit It and Quit It",
	"Super Stupid",
	"The Party",
	"Safe and Sound",
	"Runaway Trains",
	"The Damage You've Done",
	"Let Me Up (I've Had Enough)",
	"Too Much Ain't Enough",
	"Try a Little Harder",
	"Choked Up",
}

func AddFakeTracks(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		amountString := c.Query("amount")

		if amountString == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Missing query parameter \"amount\", for how many tracks to generate")
		}

		amount, err := strconv.Atoi(amountString)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("query parameter amount was not a number")
		}

		fakeTracks := []*indexFiles.IndexedTrack{}

		for i := 0; i < amount; i++ {
			fakeTracks = append(fakeTracks, generateFakeTrack())
		}

		_, err = repositories.AddTracks(fakeTracks, client, context)

		if err != nil {
			return err
		}

		return nil
	}
}

func generateFakeTrack() *indexFiles.IndexedTrack {
	artists := []string{}

	numberOfArtists := rand.Intn(2) + 1

	for i := 0; i < numberOfArtists; i++ {
		artists = append(artists, artistNames[rand.Intn(10)])
	}

	startYear := 1950

	currentYear, _, _ := time.Now().Date()

	timeSpan := currentYear - startYear

	precision := []indexFiles.ReleaseDatePrecision{
		indexFiles.PRECISION_DATE,
		indexFiles.PRECISION_MONTH,
		indexFiles.PRECISION_YEAR,
	}

	year := rand.Intn(timeSpan) + startYear
	month := rand.Intn(11) + 1
	day := rand.Intn(31) + 1

	selected_precision := precision[rand.Intn(len(precision))]

	var date *indexFiles.ReleaseDate

	switch selected_precision {
	case indexFiles.PRECISION_DATE:
		date = &indexFiles.ReleaseDate{
			Year:      year,
			Month:     month,
			Date:      day,
			Precision: selected_precision,
		}
	case indexFiles.PRECISION_YEAR:
		date = &indexFiles.ReleaseDate{
			Year:      year,
			Month:     0,
			Date:      0,
			Precision: selected_precision,
		}
	case indexFiles.PRECISION_MONTH:
		date = &indexFiles.ReleaseDate{
			Year:      year,
			Month:     month,
			Date:      0,
			Precision: selected_precision,
		}
	}

	return &indexFiles.IndexedTrack{
		Path: fakePath,
		Image: &indexFiles.Image{
			Data:     generateFakeAlbumImage(),
			MimeType: "image/png",
		},
		Artists:     artists,
		AlbumArtist: artists[rand.Intn(len(artists))],
		AlbumName:   albumNames[rand.Intn(len(albumNames))],
		Length:      rand.Intn(420) + 10,
		TrackNumber: rand.Intn(20) + 1,
		MimeType:    "audio/flac",
		Title:       trackTitles[rand.Intn(len(trackTitles))],
		Date:        date,
	}
}

// Generates a grid of a base color and shades in even chunks
func generateFakeAlbumImage() []byte {
	const imageSize = 128
	const cols = 4
	const chunks = imageSize / cols

	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	baseColor := color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: math.MaxUint8,
	}

	for col := 0; col < cols; col++ {
		for row := 0; row < cols; row++ {
			shade := color.RGBA{
				R: baseColor.R + randomOffset(),
				G: baseColor.G + randomOffset(),
				B: baseColor.B + randomOffset(),
				A: math.MaxUint8,
			}

			for x := 0; x < chunks; x++ {
				for y := 0; y < chunks; y++ {
					img.Set(
						col*chunks+x,
						row*chunks+y,
						shade)
				}
			}
		}
	}

	var buf bytes.Buffer

	png.Encode(&buf, img)

	return buf.Bytes()
}

func randomOffset() uint8 {
	return uint8(rand.Intn(24) - 12)
}
