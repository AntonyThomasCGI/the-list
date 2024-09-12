import React, { useState } from "react";


interface ShowSuggestion {
    id: number;
    title: string;
    release_date: string;
}


export const AddShow: React.FC = () => {
   const [suggestedShows, setSuggestedShows] = useState<any[]>([]);

    function handleAddEvent(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const value = event.target[0].value;
        fetch("/api/v1/search/shows?" + new URLSearchParams({"query": value}))
            .then(resp => resp.json())
            .then(data => {
                data.forEach((show: ShowSuggestion) => {
                    let year: string = show.release_date.split("-")[0];
                    console.log(`${show.title} (${year})`);
                });
                setSuggestedShows(data);
            })
            .catch(err => console.error(err));
    }

    function handleShowSelect(show: ShowSuggestion) {
        // post show?
        console.log(`Post show: ${show.title}`);
        setSuggestedShows([]);
    }

    return (
        <div>
            <div>
                <form onSubmit={handleAddEvent}>
                    <input className="add-show" type="text" placeholder="Add show"></input>
                </form>
            </div>
            <div className="dropdown">
                {suggestedShows.map((show) => {
                    let year: string = show.release_date.split("-")[0];
                    return(
                        <button key={show.id} onClick={() => handleShowSelect(show)}>{`${show.title} (${year})`}</button>
                    )
                })}
            </div>
        </div>
    )
}
