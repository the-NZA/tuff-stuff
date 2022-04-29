module.exports = (env) => {
	const plugins = [
		require("postcss-import")(),
		require("postcss-sort-media-queries")({
			sort: 'desktop-first'
		}),
		require("autoprefixer")(),
		require("postcss-preset-env")({
			stage: 2,
			features: {
				"custom-media-queries": true,
				"custom-properties": true,
				"custom-selectors": true,
				"nesting-rules": true,
				"system-ui-font-family": true
			}
		}),
		require('postcss-media-minmax')(),
	]

	if (env.mode === "production") {
		plugins.push(
			require("cssnano")({
				preset: "default",
			}),
		)
	}

	return {
		plugins
	}
}