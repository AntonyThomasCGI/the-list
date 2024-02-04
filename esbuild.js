const esbuild = require("esbuild");

esbuild.context({
    entryPoints: ["frontend/App.tsx", "frontend/App.css"],
    outdir: "public/assets",
    bundle: true,
})
.then((ctx) => {
    console.log("⚡ Build complete! ⚡");
    ctx.watch();
})
