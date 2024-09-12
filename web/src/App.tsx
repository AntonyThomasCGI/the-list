import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom/client";
import { Show, ShowComponent } from "./components/Show";
import { SearchBar } from "./components/SearchBar";


function App() {
    const [shows, setShows] = useState<Array<Show>>([]);

    const [filterTitle, setFilterTitle] = useState<string>("");
    const [filterWatched, setFilterWatched] = useState<boolean>(false);

    useEffect(() => {
    fetch("/api/v1/shows")
    .then(resp => resp.json())
    .then(data => {
        setShows(data);
    })
    .catch(err => console.log(err));
    }, []);


    function filterShow(show: Show): boolean {
        if (filterTitle && !show.title.toLowerCase().includes(filterTitle)) {
            return false;
        }
        if (filterWatched && show.watched) {
            return false;
        }
        return true;
    }

    return (
        <div>
            <header className="site-header">
                <SearchBar 
                    onFilterTitleChanged={(newValue: string) => setFilterTitle(newValue)}
                    onFilterWatchedChanged={(newValue: boolean) => setFilterWatched(newValue)}
                />
            </header>
            <div className="centered">
                <div id="list">
                    {shows.map((show, idx) => {
                        if (filterShow(show)) {
                            return <ShowComponent key={idx} id={show.id} title={show.title} author={show.author} watched={show.watched}/>
                        }
                    })}
                </div>
            </div>
        </div>
    )
}

const appDiv = document.getElementById('the-list-app')!;
const root = ReactDOM.createRoot(appDiv);
root.render(<App />);
