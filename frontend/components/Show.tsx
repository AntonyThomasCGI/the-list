
import React from "react";


interface ShowProps {
    title: string;
    author: string;
}

export const Show: React.FC<ShowProps> = (props) => {
    return (
        <div className="show">
            <h1>{props.title}</h1>
            <p>Author: {props.author}</p>
        </div>
    )
}
