
import React, { useState } from "react";


interface ShowProps {
    title: string;
    author: string;
    watched: boolean;
}

export const Show: React.FC<ShowProps> = (props) => {
    const [watched, setWatched] = useState(props.watched);

    function handleChecked(event: React.ChangeEvent<HTMLInputElement>) {
        setWatched(!watched);
        // TODO, post update to show when endpoint exists
    }

    return (
        <div className="show">
            <div className={watched ? "show-title watched" : "show-title"}>{props.title}</div>
            <div className="spacer"></div>
            <div className="checkbox-wrapper">
                <input id="watched-checkbox" type="checkbox" defaultChecked={watched} onChange={handleChecked}></input>
            </div>
        </div>
    )
}
