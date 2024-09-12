import React, { useState } from "react";

import { AddShow } from "./AddShow";
import { Show } from "./Show";


interface SearchBarProps {
    onFilterTitleChanged: (newValue: string) => void;
    onFilterWatchedChanged: (newValue: boolean) => void;
}


export const SearchBar: React.FC<SearchBarProps> = (props) => {

    return (
        <div className="topbar-container">
            <div className="site-logo">The List</div>
            {/*<AddShow />*/}
            <input className="add-show" type="text" placeholder="Search" onChange={
                (event) => props.onFilterTitleChanged(event.target.value.toLowerCase())
            }></input>
            <input type="checkbox" onChange={
                (event) => props.onFilterWatchedChanged(event.target.checked)
            }></input>
        </div>
    )
}
