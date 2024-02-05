const esbuild = require("esbuild");


const watch = process.argv.slice(2).includes("--watch");

const buildConfig = {
    entryPoints: ["frontend/App.tsx", "frontend/App.css"],
    outdir: "public/assets",
    bundle: true,
    logLevel: "info",
};

if (watch) {
    esbuild.context(buildConfig)
        .then((ctx) => {
            console.log("👀Watching for changes")
            ctx.watch();
        });
} else {
    esbuild.build(buildConfig)
        .then(() => {
            console.log("⚡ Build complete! ⚡");
        })
}

