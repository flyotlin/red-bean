import React from "react";
import '../style/menubar.css';

interface Props {
    text: string;
}

function MenuBarItem(props: Props) {
    return (
        <div className="menu-bar-item">{ props.text }</div>
    )
}

export default MenuBarItem;
