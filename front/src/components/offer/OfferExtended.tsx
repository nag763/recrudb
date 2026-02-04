import type { Offer } from "@/types/offer";
import ReactMarkdown from "react-markdown";

export default function OfferExtended({ offer }: Readonly<{offer: Offer}>) {
    return <div className="max-h-screen w-full p-10 overflow-auto max-w-4xl mx-auto bg-white shadow-lg rounded-xl border border-gray-100">
        {/* Header */}
        <div className="flex justify-between items-start border-b pb-4 mb-6">
            <div>
                <h1 className="text-2xl font-bold text-gray-800">{offer.name}</h1>
                <p className="text-gray-500">
                    {offer?.address?.city}, {offer?.address?.country} • Hybrid
                </p>
            </div>
            <div className="bg-yellow-100 text-yellow-800 px-4 py-2 rounded-full font-bold">
                Score: {offer.candidateScore}/20
            </div>
        </div>

        <div className="bg-indigo-50 p-4 rounded-lg">
            <h3 className="font-bold mb-2 text-indigo-900">Agent Recap</h3>
            <div className="text-indigo-800 text-sm whitespace-pre-line">
                <ReactMarkdown>{offer.agentRecap}</ReactMarkdown>{' '}
                {/* Shortened for clean UI */}
            </div>
        </div>

        {/* Tech Stack Tags */}
        <div className="mb-8">
            <h3 className="text-sm font-semibold uppercase tracking-wider text-gray-400 mb-3">
                Tech Requirements
            </h3>
            <div className="flex flex-wrap gap-2">
                {offer?.techRequirements?.map((tech) => (
                    <span
                        key={tech}
                        className="bg-blue-50 text-blue-700 px-3 py-1 rounded-md text-sm font-medium border border-blue-100"
                    >
                        {tech}
                    </span>
                ))}
            </div>
        </div>

        {/* Recap Section */}
        <div className="gap-8 mb-8">
            <div className="bg-gray-50 p-4 rounded-lg">
                <h3 className="font-bold mb-2 text-gray-700">Role Description</h3>
                <p className="text-gray-600 text-sm leading-relaxed">
                    {offer.description}
                </p>
            </div>
        </div>

        {/* Footer / Apply */}
        <div className="bg-teal-50 border border-teal-100 p-4 rounded-lg">
            <h3 className="text-teal-800 font-bold mb-1">How to apply?</h3>
            <p className="text-teal-700 text-sm mb-3">{offer.howToApply}</p>
        </div>
    </div>
}