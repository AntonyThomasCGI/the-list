
import React from "react";


interface ShowProps {
    title: string;
    author: string;
}

export const Show: React.FC<ShowProps> = (props) => {
    return (
        <div className="show">
            <div className="title">{props.title}</div>
            <div>Author: {props.author}</div>
        </div>
    )
}
