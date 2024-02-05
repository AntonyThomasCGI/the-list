const esbuild = require("esbuild");


const watch = process.argv.slice(2).includes("--watch");

esbuild.context({
    entryPoints: ["frontend/App.tsx", "frontend/App.css"],
    outdir: "public/assets",
    bundle: true,
})
.then((ctx) => {
    console.log("⚡ Build complete! ⚡");

    if (watch) {
        console.log("👀 Watching for changes...")
        ctx.watch();
    }
});
