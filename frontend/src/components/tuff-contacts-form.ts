import { LitElement, html, css } from 'lit';
import { property } from 'lit/decorators.js';

// @customElement("tuff-contacts-form")
export class TuffContactsForm extends LitElement {
	static styles = css`
		/* p{
			color: coral;
			max-width: 50%;
			margin: 16px auto;
		} */
	`;

	@property()
	mood = "great";

	render() {
		return html`
			<div>
				<p>This is ${this.mood} day!</p>
				<p>This is just another string</p>
			</div>
		`;
	}
}

// Register custom elements
// window.customElements.define('tuff-contacts-form', TuffContactsForm)