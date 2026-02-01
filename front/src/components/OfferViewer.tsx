import { OfferContext } from "@/context/OfferContext"
import { getOffer } from "@/interactions/offers";
import { useQuery } from "@tanstack/react-query";
import { useContext, useMemo } from "react"
import ReactMarkdown from 'react-markdown';

export default function OfferViewer() {

    const selectedOffer = useContext(OfferContext);

    const { data, isPending } = useQuery({
        queryFn: () => getOffer(selectedOffer?.id),
        queryKey: ['offer-id', selectedOffer?.id],
        enabled: !!selectedOffer?.id
    })

    const offer = useMemo(() => data?.data, [data]);

    if(!selectedOffer) {
        return <div className="w-full h-screen text-4xl text-gray-300 align-middle text-center">Select an offer</div>
    }

    if (isPending) {
        return <p>Loading...</p>
    }


    return <div className="max-h-screen w-screen p-8  overflow-auto">
        <div className="max-w-4xl mx-auto p-6 bg-white shadow-lg rounded-xl border border-gray-100">
            {/* Header */}
            <div className="flex justify-between items-start border-b pb-4 mb-6">
                <div>
                    <h1 className="text-2xl font-bold text-gray-800">{offer.name}</h1>
                    <p className="text-gray-500">{offer.address.city}, {offer.address.country} • Hybrid</p>
                </div>
                <div className="bg-yellow-100 text-yellow-800 px-4 py-2 rounded-full font-bold">
                    Score: {offer.candidateScore}/20
                </div>
            </div>

            <div className="bg-indigo-50 p-4 rounded-lg">
                <h3 className="font-bold mb-2 text-indigo-900">Agent Recap</h3>
                <div className="text-indigo-800 text-sm whitespace-pre-line">
                    <ReactMarkdown>{offer.agentRecap}</ReactMarkdown> {/* Shortened for clean UI */}
                </div>
            </div>

            {/* Tech Stack Tags */}
            <div className="mb-8">
                <h3 className="text-sm font-semibold uppercase tracking-wider text-gray-400 mb-3">Tech Requirements</h3>
                <div className="flex flex-wrap gap-2">
                    {offer.techRequirements.map((tech) => (
                        <span key={tech} className="bg-blue-50 text-blue-700 px-3 py-1 rounded-md text-sm font-medium border border-blue-100">
                            {tech}
                        </span>
                    ))}
                </div>
            </div>

            {/* Recap Section */}
            <div className="gap-8 mb-8">
                <div className="bg-gray-50 p-4 rounded-lg">
                    <h3 className="font-bold mb-2 text-gray-700">Role Description</h3>
                    <p className="text-gray-600 text-sm leading-relaxed">{offer.description}</p>
                </div>

            </div>

            {/* Footer / Apply */}
            <div className="bg-teal-50 border border-teal-100 p-4 rounded-lg">
                <h3 className="text-teal-800 font-bold mb-1">How to apply?</h3>
                <p className="text-teal-700 text-sm mb-3">{offer.howToApply}</p>
            </div>
        </div>
    </div>
}