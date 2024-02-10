const esbuild = require("esbuild");


const watch = process.argv.slice(2).includes("--watch");

const buildConfig = {
    entryPoints: ["web/src/App.tsx", "web/src/App.css"],
    outdir: "web/public/assets",
    bundle: true,
    logLevel: "info",
    loader: {
        ".png": "dataurl",
        ".woff": "dataurl",
        ".woff2": "dataurl",
        ".eot": "dataurl",
        ".ttf": "dataurl",
        ".svg": "dataurl",
    },
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

