import axios from "axios"

export const baseURL = window.location.origin

// Return new config axios with baseURL.
export default axios.create({
	baseURL: baseURL,
})