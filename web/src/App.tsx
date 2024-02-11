import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom/client";
import { Show } from "./Show";
import { AddShow } from "./AddShow";


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
        <div>
            <header className="site-header">
                <div className="site-logo">The List</div>
            </header>
            <div className="centered">
                <AddShow />
                <div id="list">
                    {shows.map((show, idx) => {
                        return <Show key={idx} id={show.id} title={show.title} author={show.author} watched={show.watched}/>
                    })}
                </div>
            </div>
        </div>
    )
}

const appDiv = document.getElementById('the-list-app')!;
const root = ReactDOM.createRoot(appDiv);
root.render(<App />);
