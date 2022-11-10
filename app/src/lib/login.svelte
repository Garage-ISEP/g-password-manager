<script lang="ts">
</script>

<header>
	<h1>Garage Vault</h1>
	<div class="imgs">
		<img src="garage.png" alt="Garage logo" />
		<img src="lock.png" alt="Lock icon" />
	</div>
</header>

<div
	id="g_id_onload"
	data-client_id="576278475266-hr36e0gt75c0u7j4pqglck2se9qeui3r.apps.googleusercontent.com"
	data-login_uri="http://192.168.1.22:5173"
	data-auto_prompt="false"
/>
<div
	class="g_id_signin g-button"
	data-type="standard"
	data-size="large"
	data-theme="outline"
	data-text="sign_in_with"
	data-shape="rectangular"
	data-logo_alignment="left"
/>

<style lang="scss">
	header {
		width: 100%;
		display: flex;
		flex-direction: column;
		margin-top: 10vh;
	}
	h1 {
		font-size: 3rem;
		text-align: center;
	}
	.imgs {
		display: flex;
		justify-content: center;
		width: 90%;
		margin: 0 auto;
		img {
			display: block;
			margin: 0 auto;
			width: 40%;
			max-width: 250px;
			margin-top: 30px;
		}
	}
	.g-button {
		display: block;
		margin: 0 auto;
		margin-top: 50px;
		width: 95%;
		max-width: 250px;
		box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.5);
	}
</style>
