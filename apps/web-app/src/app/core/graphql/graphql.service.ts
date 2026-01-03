import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GRAPHQL_ENDPOINT } from './graphql.tokens';
import { print } from 'graphql';
import type { DocumentNode } from 'graphql';
import {
  GraphQLResponse,
  GraphQLRequestError,
  GraphQLErrorCode,
} from './graphql-errors';

export type GraphQLDocument = DocumentNode;

@Injectable({
  providedIn: 'root',
})
export class GraphqlService {
  private readonly http = inject(HttpClient);
  private readonly endpoint = inject(GRAPHQL_ENDPOINT);

  /**
   * Low-level GraphQL request helper with structured error handling
   * @template TData - The expected response data type
   * @template TVariables - The variables type for the GraphQL operation
   * @template TErrorCode - Specific error codes that can occur for this operation
   */
  request<
    TData,
    TVariables,
    TErrorCode extends GraphQLErrorCode = GraphQLErrorCode,
  >(
    document: GraphQLDocument,
    options?: {
      variables?: TVariables;
      operationName?: string;
    },
  ): Observable<TData> {
    return this.http
      .post<GraphQLResponse<TData>>(this.endpoint, {
        query: print(document),
        variables: options?.variables,
        operationName: options?.operationName,
      })
      .pipe(
        map((res) => {
          if (res.errors?.length) {
            const primaryError = res.errors[0];
            console.error('GraphQL Error:', {
              code: primaryError.extensions?.code,
              message: primaryError.message,
              field: primaryError.extensions?.field,
              path: primaryError.path,
              totalErrors: res.errors.length,
            });

            throw new GraphQLRequestError<TErrorCode>(
              primaryError.message ?? 'GraphQL request failed',
              res.errors,
            );
          }

          if (res.data === undefined) {
            console.error('GraphQL Error: Response contained no data');
            throw new Error('GraphQL response contained no data');
          }

          return res.data;
        }),
      );
  }

  /**
   * Semantic wrapper for queries
   * @template TData - The expected response data type
   * @template TVariables - The variables type for the GraphQL query
   * @template TErrorCode - Specific error codes that can occur for this query
   */
  query<
    TData,
    TVariables,
    TErrorCode extends GraphQLErrorCode = GraphQLErrorCode,
  >(
    document: GraphQLDocument,
    variables?: TVariables,
    operationName?: string,
  ): Observable<TData> {
    return this.request<TData, TVariables, TErrorCode>(document, {
      variables,
      operationName,
    });
  }

  /**
   * Semantic wrapper for mutations
   * @template TData - The expected response data type
   * @template TVariables - The variables type for the GraphQL mutation
   * @template TErrorCode - Specific error codes that can occur for this mutation
   */
  mutate<
    TData,
    TVariables,
    TErrorCode extends GraphQLErrorCode = GraphQLErrorCode,
  >(
    document: GraphQLDocument,
    variables?: TVariables,
    operationName?: string,
  ): Observable<TData> {
    return this.request<TData, TVariables, TErrorCode>(document, {
      variables,
      operationName,
    });
  }
}
