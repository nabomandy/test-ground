// eslint-disable-next-line @typescript-eslint/no-var-requires,no-undef
const { createProxyMiddleware } = require('http-proxy-middleware')
// eslint-disable-next-line no-undef
module.exports = function (app) {
	
	
	app.use(
		createProxyMiddleware("/vurix-dms/api", {
		  target: "http://localhost:8080",
		  changeOrigin: true,
		})
	  );

}