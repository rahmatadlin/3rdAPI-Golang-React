export default function Card({ book, url }) {
  return (
    <>
      <div className="card bg-base-100 shadow-2xl flex flex-row">
        <figure>
          <img className="p-4 w-full" src={book.preview} alt="book image" />
        </figure>
        <div className="card-body flex-1 justify-between">
          <b className="card-title">{book.title}</b>
          <div className="flex-1">{book.author}</div>
        </div>
      </div>
    </>
  );
}
