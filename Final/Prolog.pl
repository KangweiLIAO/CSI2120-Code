weekends(Month, Year, WeekEndTemperature, Normals).
weekends(march,2020, [-4, 1, 6, 4,-2,-4, 0, 7, 8],[-1, 0, 0, 2, 2, 4, 4, 6 ,6]).

% a)
difference([],[],[]).               % Base case
difference([H|T], [H1|T1], D) :-    
    difference(T, T1, D1),          % Recursive call untill list empty
    R is H-H1,                      % Find the difference
    D = [R|D1].                     % Append result

% b)
positive([],0).             % Base case
positive([H|T], N) :-
    H>0,                    % True if the head is above 0
    positive(T,N1),         % Cummulate the counter N
    N is N1+1.              % Counter++
positive([_|T],N) :-
    positive(T,N), !.       % Cases if the head is smaller than 0

% c)
niceMonth(M,Y) :- 
    weekends(M,Y,T,N),      % Find the data(T, N) from database
    difference(T,N,D),      % Find the difference(D) of the given days
    positive(D,NN),         % Find the number(NN) of weekend days for a month that were above normal
    length(T, L),           % Find the length of given data
    NN>L/2.                 % True if at least half of the data were above normal
    