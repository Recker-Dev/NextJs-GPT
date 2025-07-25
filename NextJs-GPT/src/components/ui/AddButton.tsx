import React, { useState, useRef, useEffect } from "react";
import clsx from "clsx";
import { NotebookPen, FileStack, Plus } from 'lucide-react';

interface AddButtonProps {
    showAddMemoryModal: () => void;
    showAddFilesModal: () => void;
}


export const AddButton: React.FC<AddButtonProps> = ({
    showAddMemoryModal,
    showAddFilesModal
}) => {
    const [isOpen, setIsOpen] = useState(false);
    const buttonRef = useRef<HTMLButtonElement>(null);
    const dropdownRef = useRef<HTMLDivElement>(null);

    // Close dropdown if clicking outside
    useEffect(() => {
        function handleClickOutside(event: MouseEvent) {
            if (
                dropdownRef.current &&
                !dropdownRef.current.contains(event.target as Node) &&
                !buttonRef.current?.contains(event.target as Node)
            ) {
                setIsOpen(false);
            }
        }

        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);

    return (
        <div className="relative">
            {/* PLUS button */}
            <button
                ref={buttonRef}
                onClick={() => setIsOpen(!isOpen)}
                className={clsx(
                    "flex items-center justify-center",
                    "w-12 h-12 rounded-full bg-gradient-to-r from-green-600 to-green-700",
                    "hover:from-green-500 hover:to-green-600 hover:scale-105",
                    "transition-all duration-150 ease-in-out",
                    "text-white"
                )}
                aria-label="Add Button"
            >
                <Plus />
            </button>

            {/* Dropdown menu */}
            {isOpen && (
                <div
                    ref={dropdownRef}
                    className={clsx(
                        "absolute left-0 bottom-full mb-2 z-50",
                        "flex flex-col gap-2 w-48 rounded-lg shadow-lg",
                        "bg-gray-800 border border-gray-700"
                    )} >
                    <div
                        className={clsx("flex items-center gap-2 px-2  hover:bg-gray-700 hover:shadow-xl my-1 mx-1 rounded-lg transition-all",
                        )}
                    >
                        {/* <svg
                            xmlns="http://www.w3.org/2000/svg"
                            className="h-6 w-6"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke="currentColor"
                            strokeWidth="2"
                        >
                            <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                d="M5 19V7.5a2.5 2.5 0 012.5-2.5h9A2.5 2.5 0 0119 7.5V19m-9 0v-4a2 2 0 00-2-2H8a2 2 0 00-2 2v4m3-11H8m6 0h-3"
                            />
                        </svg> */}
                        <NotebookPen className="text-purple-400" />
                        <button
                            className={clsx(
                                "flex-1 w-full h-full", 
                                "text-left py-2 text-white",
                                "transition-colors duration-150",
                                "hover:bg-gradient-to-r hover:from-purple-400 hover:to-pink-600 hover:bg-clip-text hover:text-transparent"
                            )}
                            onClick={() => {
                                console.log("Add Memory clicked!");
                                showAddMemoryModal();
                                setIsOpen(false);
                            }}
                        >
                            Add Memory
                        </button>
                    </div>
                    <div
                        className={clsx("flex items-center gap-2 px-2  hover:bg-gray-700 hover:shadow-xl my-1 mx-1 rounded-lg transition-all",
                        )}
                    >
                        <FileStack className="text-blue-400" />
                        <button
                            className={clsx(
                                "flex-1 w-full h-full", 
                                "text-left py-2 text-white",
                                "transition-colors duration-150",
                                "hover:bg-gradient-to-r hover:from-blue-400 hover:to-cyan-600 hover:bg-clip-text hover:text-transparent"

                            )}
                            onClick={() => {
                                console.log("Add Memory clicked!");
                                showAddFilesModal();
                                setIsOpen(false);
                            }}
                        >
                            Add Files
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
}
