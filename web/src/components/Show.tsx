
import React, { useState } from "react";


export interface Show {
    id: number;
    title: string;
    author: string;
    watched: boolean;
}


export const ShowComponent: React.FC<Show> = (props) => {
    const [watched, setWatched] = useState(props.watched);

    function handleChecked(event: React.ChangeEvent<HTMLInputElement>) {
        let newWatched = !watched;
        setWatched(newWatched);

        let jsonData = {
            "watched": newWatched,
        };
        fetch(`/api/v1/show/${props.id}`, {
            method: "PUT",
            body: JSON.stringify(jsonData)
        });
    }

    return (
        <div className="show">
            <div className={watched ? "show-title watched" : "show-title"}>{props.title}</div>
            <div className="spacer"></div>
            <div className="checkbox-wrapper">
                <input
                    id="watched-checkbox"
                    type="checkbox"
                    defaultChecked={watched}
                    onChange={handleChecked}>
                </input>
            </div>
        </div>
    )
}
