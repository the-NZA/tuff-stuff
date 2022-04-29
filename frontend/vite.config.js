import { defineConfig } from "vite"
import { resolve, extname } from "path"
import minifyHTML from "rollup-plugin-minify-html-literals"
import copy from "rollup-plugin-copy"

// Needed directories paths
const rootPath = resolve(__dirname, "src")
const distPath = resolve(__dirname, "dist")
const publicPath = resolve(__dirname, "public")

// Supported file types
// ref: https://github.com/vitejs/vite/blob/main/packages/vite/src/node/constants.ts
const assetsTypes = {
	styles: ['.css'],
	images: [
		'.png', '.jpg', '.jpeg', '.gif',
		'.svg', '.ico', '.webp', '.avif',
	],
	media: [
		'.mp4', '.webm', '.ogg', '.mp3',
		'.wav', '.flac', '.aac',
	],
	fonts: [
		'.woff', '.woff2', '.eot',
		'.ttf', '.otf',
	]
}

// Return output path for a given pattern
const outputPath = (pattern) => `static/${pattern}`

export default defineConfig({
	root: rootPath,
	publicDir: publicPath,
	build: {
		outDir: distPath,
		emptyOutDir: true,
		assetsDir: "assets",
		assetsInlineLimit: 0,
		cssCodeSplit: false,

		rollupOptions: {
			plugins: [
				minifyHTML(),
				copy({
					hook: "buildEnd",
					targets: [
						// Copy all static files to project root/static
						{ src: "dist/static/", dest: "../" },
					]
				})
			],
			input: {
				index: resolve(rootPath, "index.html"),
				// about: resolve(rootPath, "about.html"),
			},
			output: {
				manualChunks: undefined, // disable js chunks
				chunkFileNames: outputPath("scripts/[name].[hash].js"),
				entryFileNames: outputPath("scripts/[name].[hash].js"),
				assetFileNames: ({ name }) => {
					// ref: https://github.com/vitejs/vite/discussions/6552

					const assetExtname = extname(name)

					// If CSS file then save to dist
					if (assetsTypes.styles.includes(assetExtname)) {
						return outputPath('styles/[name].[hash][extname]')
					}

					// If any of images types then save to images
					if (assetsTypes.images.includes(assetExtname)) {
						return outputPath('images/[name][extname]')
					}

					// If any of media types then save to media
					if (assetsTypes.media.includes(assetExtname)) {
						return outputPath('media/[name][extname]')
					}

					// If any of fonts types then save to fonts
					if (assetsTypes.fonts.includes(assetExtname)) {
						return outputPath('fonts/[name][extname]')
					}

					// default value just save to assets
					// ref: https://rollupjs.org/guide/en/#outputassetfilenames
					return outputPath('assets/[name]-[hash][extname]')
				}
			}
		}
	}
})