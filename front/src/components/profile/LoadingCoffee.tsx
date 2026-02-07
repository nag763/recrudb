
export default function LoadingCoffee() {

    return (
        <div className='h-full flex gap-8 w-full items-center justify-between'>
            <img src='/coffee.svg'></img>
            <div className="flex flex-col gap-4 text-right">
                <h1 className="text-4xl animate-pulse">Your CV is being processed</h1>
                <h2 className="text-2xl">You can go for a coffee, but you will have to be fast ...</h2>
            </div>
        </div>
    );
}
