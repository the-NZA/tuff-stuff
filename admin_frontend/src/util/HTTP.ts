import axios from "axios"

// export const baseURL = window.location.origin
export const baseURL = "http://tuff-stuff.local"

// Return new config axios with baseURL.
export default axios.create({
	baseURL: baseURL,
})