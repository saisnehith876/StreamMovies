import Movie from '../movie/Movie'

const Movies = ({movies, updateMovieReview, message}) => {
    const safeMovies = Array.isArray(movies) ? movies : [];

    return (
        <div className="container mt-4">
            <div className="row">
                {safeMovies.length > 0
                    ? safeMovies.map((movie) => (
                        <Movie key={movie._id} updateMovieReview={updateMovieReview} movie={movie} />
                    ))
                    : <h2>{message}</h2>
                }
            </div>
        </div>
    )
}
export default Movies;