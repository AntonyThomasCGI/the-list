import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom/client";
import { Show } from "./components/Show";


function App() {
    const [shows, setShows] = useState<any[]>([]);

    useEffect(() => {
        fetch("/api/v1/shows")
        .then(resp => resp.json())
        .then(data => {
            setShows(data);
        })
        .catch(err => console.log(err));
    }, []);

    return (
        <div className="centered">
            <div id="list">
                {shows.map((show, idx) => {
                    return <Show key={idx} title={show.title} author={show.author}/>
                })}
            </div>
        </div>
    )
}

const appDiv = document.getElementById('the-list-app')!;
const root = ReactDOM.createRoot(appDiv);
root.render(<App />);
