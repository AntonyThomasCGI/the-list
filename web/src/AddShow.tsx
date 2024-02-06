import React from "react";


export const AddShow: React.FC = () => {

    function handleAddEvent(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const value = event.target[0].value;
        console.log(value);
        fetch("/api/v1/search/shows?" + new URLSearchParams({"query": value}))
        .then(resp => resp.json())
        .then(data => {
            console.log(data);
        })
        .catch(err => console.log(err));
    }

    return (
        <div>
            <form onSubmit={handleAddEvent}>
                <input className="add-show" type="text" placeholder="Add show"></input>
            </form>
        </div>
    )
}
