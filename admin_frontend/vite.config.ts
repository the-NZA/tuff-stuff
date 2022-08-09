import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import copy from "rollup-plugin-copy"

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [vue()],
	base: "/admin/",
	build: {
		rollupOptions: {
			plugins: [
				copy({
					hook: "buildEnd",
					targets: [
						// Copy all static files to project root/admin_static
						{src: "dist/*", dest: "../admin_static/"},
					]
				}),
			],
		}
	}
})
